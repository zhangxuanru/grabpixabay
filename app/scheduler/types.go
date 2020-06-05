/*
@Time : 2020/6/3 17:44
@Author : zxr
@File : types
@Software: GoLand
*/
package scheduler

import (
	"context"
	"grabpixabay/app/storage"
)

type Item struct {
	//inImageChan   chan *ImageInfo //
	inColorChan   chan *storage.ImgColor
	imageListChan chan *storage.ImageInfo //图片信息chan
	downloadChan  chan *storage.ImageInfo //执行下载图片
}

type Concurrent struct {
	workerCount int
	WorkActive  bool //worker状态，true 表示已启动
	Ctx         context.Context
	Cancel      context.CancelFunc
	Item
}
