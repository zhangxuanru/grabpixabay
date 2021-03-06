/*
@Time : 2020/6/15 16:49
@Author : zxr
@File : imagestorage
@Software: GoLand
*/
package services

import (
	"bytes"
	"grabpixabay/core/api"
	"grabpixabay/core/storage/models"
	"grabpixabay/core/storage/services/es"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type PicAttr struct {
	File   string
	PicId  int
	Width  int
	Height int
	Size   int
}

//保存所有信息
func (i *ImageService) SaveAll(item api.ItemImage) {
	i.SaveAuthor(item)   //保存作者信息
	i.SaveUserStat(item) //修改用户统计表
	i.SavePicture(item)  //保存图片主信息
	i.SaveTag(item)      //保存tag信息
	i.DownloadPic(item)  //下载图片保存图片属性
	i.SavePicApi(item)   //保存返回的API信息
	if _, ok := i.InsertPicMap[item.ID]; ok {
		es.SavePicInfo(item) //保存到ES中
	}
}

//保存作者信息
func (i *ImageService) SaveAuthor(item api.ItemImage) (id int) {
	if id, ok := i.AuthorMap[item.UserID]; ok {
		logrus.Println("UID ", item.UserID, "已存在")
		return id
	}
	qiNiu := &QiNiu{
		SrcFile: item.UserImageURL,
	}
	ret, err := qiNiu.UploadFile()
	if err != nil {
		log := &models.PictureAttrLog{
			PicId:    uint(item.ID),
			ImageURL: qiNiu.SrcFile,
			ErrMsg:   err.Error(),
			AddTime:  time.Now(),
		}
		_, _ = log.Insert()
	}
	isUpload := 0
	if ret != nil && ret.PutRet != nil && ret.PutRet.Key != "" {
		isUpload = 1
	}
	user := &models.User{
		PxUid:        int64(item.UserID),
		NickName:     item.User,
		UserType:     models.UserPx,
		HeadPortrait: item.UserImageURL,
		FileName:     ret.FileName,
		IsQiniu:      isUpload,
		AddTime:      time.Now(),
		UpdateTime:   time.Now(),
	}
	if id, err := user.InsertCheckByUid(); err != nil || id < 1 {
		logrus.WithFields(logrus.Fields{
			"method": "SaveAuthor",
			"data":   *user,
		}).Debug()
		logrus.Error("SaveAuthor error:", err)
	} else {
		i.AuthorMap[item.UserID] = id
	}
	return id
}

//保存图片主信息
func (i *ImageService) SavePicture(item api.ItemImage) {
	if _, ok := i.PicMap[item.ID]; ok {
		//todo 以后可加 更新数据库的 逻辑
		logrus.Println("PID ", item.ID, "已存在")
		return
	}
	pic := &models.Picture{}
	uid := 0
	if uid = i.GetUidByAuthorId(item.UserID); uid == 0 {
		uid = i.SaveAuthor(item)
	}
	if item.EditorsChoice == true {
		pic.IsHandpick = 1
	}
	if item.Orientation != "" {
		pic.ImageDirection = i.GetDirection(item.Orientation)
	}
	if item.Color != "" {
		pic.ThemeColor = item.Color
	}
	if item.Category != "" {
		if catId, ok := i.CategoryMap[item.Category]; !ok {
			catId = models.NewCategory().GetIdByCateName(item.Category)
			i.CategoryMap[item.Category] = catId
			pic.CategoryId = uint(catId)
		} else {
			pic.CategoryId = uint(catId)
		}
	}
	pic.Uid = int64(uid)
	pic.PxUid = int64(item.UserID)
	pic.ImageFormat = i.GetImageFormat(item.LargeImageURL)
	pic.ImageType = i.GetImageType(item.Type)
	pic.PxImgId = uint(item.ID)
	pic.ViewNum = uint(item.Views)
	pic.PageURL = item.PageURL
	pic.DownloadsNum = uint(item.Downloads)
	pic.LikeNum = uint(item.Likes)
	pic.FavoritesNum = uint(item.Favorites)
	pic.CommentsNum = uint(item.Comments)
	pic.State = models.StatusDefault
	pic.AddTime = time.Now()
	pic.UpdateTime = time.Now()
	if id, isCreate, err := pic.Save(); err != nil || id < 1 {
		logrus.WithFields(logrus.Fields{
			"method": "SavePicture",
			"data":   *pic,
		}).Debug()
		logrus.Error("SavePicture error:", err)
	} else {
		i.PicMap[item.ID] = id
		if isCreate == true {
			i.InsertPicMap[item.ID] = struct{}{}
			logrus.Println(item.ID, "插入数据库成功，已成功插入", len(i.InsertPicMap), "条数据")
		}
	}
}

//修改用户统计表
func (i *ImageService) SaveUserStat(item api.ItemImage) {
	uid := 0
	if uid = i.GetUidByAuthorId(item.UserID); uid == 0 {
		uid = i.SaveAuthor(item)
	}
	stat := &models.UserStat{
		Uid:          int64(uid),
		PxUid:        int64(item.UserID),
		PicNum:       1,
		ViewNum:      uint(item.Views),
		DownloadsNum: uint(item.Downloads),
		LikeNum:      uint(item.Likes),
		CommentNum:   uint(item.Comments),
		FollowerNum:  uint(item.Favorites),
		AddTime:      time.Now(),
		UpdateTime:   time.Now(),
	}
	if _, ok := i.UserStatMap[item.UserID]; !ok {
		if id := stat.GetIdByUid(uid); id > 0 {
			i.UserStatMap[item.UserID] = id
		}
	}
	if _, ok := i.UserStatMap[item.UserID]; ok {
		_, err := stat.UpdateStat()
		if err != nil {
			logrus.Error("UpdateStat err:", err)
		}
	} else {
		if id, err := stat.Insert(); err == nil && id > 0 {
			i.UserStatMap[item.UserID] = id
		}
	}
}

//保存tag信息
func (i *ImageService) SaveTag(item api.ItemImage) {
	var tagIdBuf bytes.Buffer
	tags := strings.TrimSpace(item.Tags)
	if _, ok := i.PicTagMap[item.ID]; ok || tags == "" {
		logrus.Infof("pic_id=%d   TAG已存在", item.ID)
		return
	}
	tagList := strings.Split(tags, ",")
	for _, tag := range tagList {
		if id, ok := i.TagMap[tag]; ok {
			tagIdBuf.WriteString(strconv.Itoa(id) + ",")
			continue
		}
		tagModel := models.NewTag()
		tagModel.TagName = tag
		tagModel.State = models.StatusDefault
		tagModel.AddTime = time.Now()
		tagModel.UpdateTime = time.Now()
		if id, err := tagModel.Insert(); id > 0 {
			tmp := tag
			i.TagMap[tmp] = id
			tagIdBuf.WriteString(strconv.Itoa(id) + ",")
		} else {
			logrus.Error("tagModel.Insert error :", err)
		}
	}
	//保存图片tag信息
	picTag := models.NewPictureTag()
	picTag.PicId = uint(item.ID)
	picTag.TagId = strings.TrimRight(tagIdBuf.String(), ",")
	picTag.State = models.StatusDefault
	picTag.AddTime = time.Now()
	picTag.UpdateTime = time.Now()
	if id, err := picTag.Insert(); id > 0 {
		i.PicTagMap[item.ID] = id
	} else {
		logrus.Error("picTag.Insert error :", err)
	}
}

//保存API信息
func (i *ImageService) SavePicApi(item api.ItemImage) {
	if _, ok := i.ApiMap[item.ID]; ok {
		return
	}
	apiData, err := i.Json.Marshal(item)
	if err != nil {
		return
	}
	picApi := &models.PicApi{
		PxImgId: uint(item.ID),
		Api:     string(apiData),
		AddTime: time.Now(),
	}
	if id, _ := picApi.Save(); id > 0 {
		i.ApiMap[item.ID] = struct{}{}
	}
}
