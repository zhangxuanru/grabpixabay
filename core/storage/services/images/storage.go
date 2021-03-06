package services

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"grabpixabay/configs"
	"grabpixabay/core/api"
	"grabpixabay/core/storage/models"

	"github.com/sirupsen/logrus"
)

//图片服务
type ImageService struct {
	ItemListChan chan api.ItemImage
	CloseChan    chan bool
	Ctx          context.Context
	Can          context.CancelFunc
	Json         jsoniter.API
	ServiceModels
	MapCache
}

//缓存MAP
type MapCache struct {
	AuthorMap    map[int]int
	PicMap       map[int]int
	UserStatMap  map[int]int
	TagMap       map[string]int
	PicTagMap    map[int]int
	CategoryMap  map[string]int
	ApiMap       map[int]struct{}
	InsertPicMap map[int]struct{}
}

//需要用到的模型
type ServiceModels struct {
	UserModel *models.User
	PicModel  *models.Picture
}

func NewImageService() *ImageService {
	ctx, cancel := context.WithCancel(context.TODO())
	return &ImageService{
		MapCache: MapCache{
			AuthorMap:    make(map[int]int),
			PicMap:       make(map[int]int),
			UserStatMap:  make(map[int]int),
			TagMap:       make(map[string]int),
			PicTagMap:    make(map[int]int),
			CategoryMap:  make(map[string]int),
			ApiMap:       make(map[int]struct{}),
			InsertPicMap: make(map[int]struct{}),
		},
		ServiceModels: ServiceModels{
			UserModel: models.NewUser(),
			PicModel:  models.NewPicture(),
		},
		Json:         jsoniter.ConfigCompatibleWithStandardLibrary,
		ItemListChan: make(chan api.ItemImage, configs.GConf.ItemQueueMaxLimit),
		CloseChan:    make(chan bool),
		Ctx:          ctx,
		Can:          cancel,
	}
}

//运行多个gorouting 来保存数据
func (i *ImageService) RunStorage(endChan chan bool) {
	for j := 0; j < 5; j++ {
		i.Storage(endChan)
	}
}

//存储图片信息
func (i *ImageService) Storage(endChan chan bool) {
	go func() {
		for {
			select {
			case item := <-i.ItemListChan:
				i.SaveAll(item)
				endChan <- true
			case <-i.CloseChan:
				i.Can()
				goto End
			case <-i.Ctx.Done():
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
