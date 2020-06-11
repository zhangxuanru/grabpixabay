/*
@Time : 2020/6/8 15:44
@Author : zxr
@File : types
@Software: GoLand
*/
package scheduler

import (
	"context"
	"grabpixabay/initialize"
	"os"
	"sync"
)

type Item struct {
	Command  *initialize.CommandLine
	Pool     *Concurrent
	Ctx      context.Context
	Can      context.CancelFunc
	SignChan chan os.Signal
}

//调度器结构体
type Concurrent struct {
	workerCount   int  //worker的个数
	WorkActive    bool //worker状态，true 表示已启动
	Ctx           context.Context
	Can           context.CancelFunc
	itemImageChan chan ItemImage  //图片信息
	itemVideoChan chan *ItemVideo //视频信息
	Wg            *sync.WaitGroup
}

//API请求图片接口返回信息
type ApiImageResp struct {
	Total     int         `json:"total"`
	TotalHits int         `json:"totalHits"`
	Hits      []ItemImage `json:"hits"`
}

//图片解析信息item
type ItemImage struct {
	ID              int    `json:"id"`
	PageURL         string `json:"pageURL"`
	Type            string `json:"type"`
	Tags            string `json:"tags"`
	PreviewURL      string `json:"previewURL"`
	PreviewWidth    int    `json:"previewWidth"`
	PreviewHeight   int    `json:"previewHeight"`
	WebformatURL    string `json:"webformatURL"`
	WebformatWidth  int    `json:"webformatWidth"`
	WebformatHeight int    `json:"webformatHeight"`
	LargeImageURL   string `json:"largeImageURL"`
	ImageWidth      int    `json:"imageWidth"`
	ImageHeight     int    `json:"imageHeight"`
	ImageSize       int    `json:"imageSize"`
	ItemImageStat
	ItemAuthor
}

//作者信息
type ItemAuthor struct {
	UserID       int    `json:"user_id"`
	User         string `json:"user"`
	UserImageURL string `json:"userImageURL"`
}

type ItemImageStat struct {
	Views     int `json:"views"`
	Downloads int `json:"downloads"`
	Favorites int `json:"favorites"`
	Likes     int `json:"likes"`
	Comments  int `json:"comments"`
}

//视频信息 item
type ItemVideo struct {
}
