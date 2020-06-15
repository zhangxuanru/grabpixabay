/*
@Time : 2020/6/8 15:43
@Author : zxr
@File : pool
@Software: GoLand
*/
package scheduler

import (
	"context"
	"fmt"
	"grabpixabay/configs"
	"grabpixabay/core/api"
	"grabpixabay/core/storage/services"
	"sync"

	"github.com/sirupsen/logrus"
)

func NewConcurrent(workerCount int, ctx context.Context, can context.CancelFunc) *Concurrent {
	return &Concurrent{
		itemImageChan: make(chan api.ItemImage),
		itemVideoChan: make(chan *api.ItemVideo),
		ItemEndChan:   make(chan bool, configs.GConf.ItemQueueMaxLimit),
		workerCount:   workerCount,
		Ctx:           ctx,
		Can:           can,
		Wg:            &sync.WaitGroup{},
		ImageStorage:  services.NewImageService(),
	}
}

func (c *Concurrent) WorkerRun() {
	if c.WorkActive == true {
		logrus.Infoln("worker 正在运行中.....")
		return
	}
	for i := 0; i < c.workerCount; i++ {
		c.createWorker(i)
	}
	c.WorkActive = true
}

//创建工作进程
func (c *Concurrent) createWorker(i int) {
	go func() {
		for {
			select {
			case image := <-c.itemImageChan:
				c.ImageStorage.AddQueueItem(image)
			case video := <-c.itemVideoChan:
				logrus.Printf("video %+v\n\n", video)
			case <-c.ItemEndChan: //数据库处理完，这里任务减一
				c.Done()
			case <-c.Ctx.Done():
				fmt.Println("Worker", i, "终止请求.....")
				return
			}
		}
	}()
}

//分发图片item
func (c *Concurrent) DistributeImageItem(resp *api.ApiImageResp) {
	if len(resp.Hits) == 0 {
		return
	}
	c.AddWgNum(len(resp.Hits))
	for _, image := range resp.Hits {
		c.SubmitImageItem(image)
	}
}

//发送图片源信息
func (c *Concurrent) SubmitImageItem(item api.ItemImage) {
	c.itemImageChan <- item
}

func (c *Concurrent) Wait() {
	c.Wg.Wait()
}

func (c *Concurrent) AddWgNum(n int) {
	c.Wg.Add(n)
}

func (c *Concurrent) Done() {
	c.Wg.Done()
}
