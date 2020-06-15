package services

import (
	"grabpixabay/configs"
	"grabpixabay/core/api"
	"grabpixabay/core/storage/models"

	"github.com/sirupsen/logrus"
)

//图片服务
type ImageService struct {
	AuthorMap    map[int]int
	PicMap       map[int]int
	UserStatMap  map[int]int
	TagMap       map[string]int
	UserModel    *models.User
	PicModel     *models.Picture
	ItemListChan chan api.ItemImage
	CloseChan    chan bool
}

func NewImageService() *ImageService {
	return &ImageService{
		AuthorMap:    make(map[int]int),
		PicMap:       make(map[int]int),
		UserStatMap:  make(map[int]int),
		TagMap:       make(map[string]int),
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
				i.SaveAll(item)
				endChan <- true
			case <-i.CloseChan:
				goto End
			}
		}
	End:
		logrus.Println("Storage service close.....")
		return
	}()
}

//关闭存储服务进程
func (i *ImageService) Close() {
	i.CloseChan <- true
}

//添加item到队列
func (i *ImageService) AddQueueItem(item api.ItemImage) {
	i.ItemListChan <- item
}
