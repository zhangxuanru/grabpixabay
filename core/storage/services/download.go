package services

import (
	"github.com/sirupsen/logrus"
	"grabpixabay/core/api"
	"grabpixabay/core/storage/models"
	"sync"
	"time"
)

//下载图片
func (i *ImageService) DownloadPic(item api.ItemImage) {
	wg := &sync.WaitGroup{}
	list := i.picAttrList(item)
	wg.Add(len(list))
	for _, attr := range list {
		attrCopy := attr
		go func(attr *PicAttr) {
			defer wg.Done()
			if attr.File == "" {
				return
			}
			i.SaveDbAttr(attr)
		}(attrCopy)
	}
	wg.Wait()
}

//保存图片属性到表中
func (i *ImageService) SaveDbAttr(attr *PicAttr) {
	pictureAttr := &models.PictureAttr{
		PicId: uint(attr.PicId),
		Width: attr.Width,
	}
	pic := pictureAttr.GetIdByPicId()
	if pic != nil && pic.IsQiniu == 1 {
		return
	}
	tmp := attr
	qiNiu := &QiNiu{
		SrcFile: tmp.File,
	}
	logrus.Println("开始下载图片:", qiNiu.SrcFile)
	ret, err := qiNiu.UploadFile()
	logrus.Println("下载图片结束:", qiNiu.SrcFile, "err:", err)
	if err != nil {
		log := &models.PictureAttrLog{
			PicId:    uint(attr.PicId),
			ImageURL: attr.File,
			ErrMsg:   err.Error(),
			AddTime:  time.Now(),
		}
		_, _ = log.Insert()
	}
	isUpload := 0
	if ret.PutRet != nil && ret.PutRet.Key != "" {
		isUpload = 1
	}
	pictureAttr = &models.PictureAttr{
		PicId:    uint(attr.PicId),
		ImageURL: attr.File,
		Width:    attr.Width,
		Height:   attr.Height,
		FileName: ret.FileName,
		IsQiniu:  isUpload,
		State:    models.StatusDefault,
		AddTime:  time.Now(),
	}
	if pic == nil || pic.Id == 0 {
		if _, err := pictureAttr.Insert(); err != nil {
			logrus.Error("pictureAttr.Insert error:", err)
		}
		return
	}
	if pic != nil && pic.Id > 0 && isUpload == 1 {
		pictureAttr.Id = pic.Id
		_, _ = pictureAttr.EditUpload(isUpload)
	}
}

//要下载的图片列表
func (i *ImageService) picAttrList(item api.ItemImage) (list []*PicAttr) {
	list = []*PicAttr{
		{
			File:   item.PreviewURL,
			PicId:  item.ID,
			Width:  item.PreviewWidth,
			Height: item.PreviewHeight,
			Size:   0,
		},
		{
			File:   item.LargeImageURL,
			PicId:  item.ID,
			Width:  960,
			Height: 1280,
			Size:   0,
		},
		{
			File:   "",
			PicId:  item.ID,
			Width:  486,
			Height: 340,
			Size:   0,
		},
		{
			File:   "",
			PicId:  item.ID,
			Width:  686,
			Height: 480,
			Size:   0,
		},
		{
			File:   "",
			PicId:  item.ID,
			Width:  960,
			Height: 720,
			Size:   0,
		},
	}
	return
}
