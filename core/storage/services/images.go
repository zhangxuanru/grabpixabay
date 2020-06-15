package services

import (
	"fmt"
	"grabpixabay/configs"
	"grabpixabay/core/api"
	"grabpixabay/core/storage/models"
	"time"

	"github.com/sirupsen/logrus"
)

//图片服务
type ImageService struct {
	AuthorMap    map[int]int
	PicMap       map[int]int
	UserModel    *models.User
	PicModel     *models.Picture
	ItemListChan chan api.ItemImage
	CloseChan    chan bool
}

func NewImageService() *ImageService {
	return &ImageService{
		AuthorMap:    make(map[int]int),
		PicMap:       make(map[int]int),
		UserModel:    models.NewUser(),
		PicModel:     models.NewPicture(),
		ItemListChan: make(chan api.ItemImage, configs.GConf.ItemQueueMaxLimit),
		CloseChan:    make(chan bool),
	}
}

//存储图片信息 //todo 这里重写， 用队列的思想来实现存储
func (i *ImageService) Storage(endChan chan bool) {
	go func() {
		for {
			select {
			case item := <-i.ItemListChan:
				fmt.Println("item:", item)
				endChan <- true
			//i.SaveAuthor(item)
			case <-i.CloseChan:
				goto End
			}
		}
	End:
		logrus.Println("Storage service close.....")
		return
	}()

	//i.SavePicture(&item)
	//fmt.Printf("storage images:%+v\n", item)
}

//关闭存储服务
func (i *ImageService) Close() {
	i.CloseChan <- true
}

func (i *ImageService) AddQueueItem(item api.ItemImage) {
	i.ItemListChan <- item
}

//保存作者信息
func (i *ImageService) SaveAuthor(item api.ItemImage) (id int) {
	fmt.Printf("map: %+v", i.AuthorMap)
	fmt.Println()
	if id, ok := i.AuthorMap[item.UserID]; ok {
		logrus.Println("UID ", item.UserID, "已存在")
		return id
	}
	user := &models.User{}
	user.PxUid = int64(item.UserID)
	user.NickName = item.User
	user.UserType = models.UserPx
	user.AddTime = time.Now()
	user.UpdateTime = time.Now()
	user.HeadPortrait = item.UserImageURL

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

//保存图片信息
func (i *ImageService) SavePicture(item api.ItemImage) {
	if _, ok := i.PicMap[item.ID]; ok {
		logrus.Println("PID ", item.ID, "已存在")
		return
	}
	pic := &models.Picture{}
	uid := 0
	if uid = i.GetUidByAuthorId(item.UserID); uid == 0 {
		uid = i.SaveAuthor(item)
	}
	pic.Uid = int64(uid)
	pic.PxUid = int64(item.UserID)
	pic.ImageFormat = GetImageFormat(item.LargeImageURL)
	pic.ImageType = GetImageType(item.Type)
	pic.PxImgId = uint(item.ID)
	pic.ViewNum = uint(item.Views)
	pic.PageURL = item.PageURL
	pic.DownloadsNum = uint(item.Downloads)
	pic.LikeNum = uint(item.Likes)
	pic.FavoritesNum = uint(item.Favorites)
	pic.CommentsNum = uint(item.Comments)
	pic.AddTime = time.Now()
	pic.UpdateTime = time.Now()
	pic.State = 1

	fmt.Printf("test PIC :%+v", pic)
	fmt.Println()
	return

	if id, err := pic.Save(); err != nil || id < 1 {
		logrus.Error("pic save error:", err)
	} else {
		i.PicMap[item.ID] = id
	}
	//修改用户统计表---todo,并发插入用户表还有问题
}

//根据图片作者ID获取真实的用户ID
func (i *ImageService) GetUidByAuthorId(authorId int) int {
	if id, ok := i.AuthorMap[authorId]; ok {
		return id
	}
	if id := models.NewUser().GetUidByAuthorId(authorId); id > 0 {
		return id
	}
	return 0
}
