/*
@Time : 2020/6/8 15:44
@Author : zxr
@File : types
@Software: GoLand
*/
package scheduler

import (
	"context"
	"grabpixabay/core/api"
	"grabpixabay/core/storage/services/images"
	"grabpixabay/initialize"
	"os"
	"sync"
)

type Task struct {
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
	itemImageChan chan api.ItemImage  //图片信息
	itemVideoChan chan *api.ItemVideo //视频信息
	Wg            *sync.WaitGroup
	ItemEndChan   chan bool
	ImageStorage  *services.ImageService
}
