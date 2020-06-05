/*
@Time : 2020/6/3 16:35
@Author : zxr
@File : scheduler
@Software: GoLand
*/
//调度images结构体
package scheduler

import (
	"fmt"
	"grabpixabay/app/storage"
	"grabpixabay/config"
	"sync"
)

var Scheduler *Concurrent
var once sync.Once

//调度器接口
type SchedulingPool interface {
	SubmitImage(*storage.ImageInfo)
	SubmitColor(*storage.ImgColor)
}

func NewConcurrent(workCount int) *Concurrent {
	once.Do(func() {
		Scheduler = &Concurrent{
			workerCount: workCount,
			Item: Item{
				inColorChan:   make(chan *storage.ImgColor),
				imageListChan: make(chan *storage.ImageInfo, config.GConf.MaxImageListSize),
				downloadChan:  make(chan *storage.ImageInfo),
			},
		}
	})
	return Scheduler
}

//调度器开始执行
func (c *Concurrent) Run() {
	if c.WorkActive == true {
		return
	}
	for i := 0; i <= c.workerCount; i++ {
		c.createWorker(i)
	}
	c.WorkActive = true
	go func() {
		for image := range c.imageListChan {
			c.downloadChan <- image
		}
	}()
}

//创建工作线程
func (c *Concurrent) createWorker(i int) {
	go func() {
		for {
			select {
			case color := <-c.inColorChan:
				fmt.Println("go color:", i, ">>", color)
			case image := <-c.downloadChan:
				storage.NewStorageImage(image).Storage()
			case <-c.Ctx.Done():
				fmt.Println("Worker", i, "终止请求.....")
				return
			}
		}
	}()
}

func (c *Concurrent) SubmitImage(image *storage.ImageInfo) {
	c.imageListChan <- image
}

func (c *Concurrent) SubmitColor(color *storage.ImgColor) {
	c.inColorChan <- color
}
