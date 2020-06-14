package services

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"grabpixabay/configs"
	"grabpixabay/core/api"
	"grabpixabay/core/storage/models"
	"time"
)

//图片服务
type ImageService struct {
	AuthorMap map[int]struct{}
}

func NewImageService() *ImageService {
	return &ImageService{
		AuthorMap: make(map[int]struct{}),
	}
}

//存储图片信息
func (i *ImageService) Storage(resp api.ItemImage) {
	fmt.Printf("storage images:%+v\n", resp)
}

//保存作者信息
func (i *ImageService) SaveAuthor() {
	user := models.NewUser()

	user.AddTime = time.Now()
	user.UpdateTime = time.Now()
	user.HeadPortrait = ""
	user.Uid = 1
	user.PersonAddr = ""
	user.NickName = ""
	user.UserType = configs.UserTypePx

	if _, err := user.InsertCheckByUid(); err != nil {
		logrus.WithFields(logrus.Fields{
			"method": "SaveAuthor",
			"data":   *user,
		}).Error(err)
	}
	return
}
