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
	"grabpixabay/common/qiniu"
	"sync"
)

var Scheduler *Concurrent
var once sync.Once

const ImageChanSize = 50000

//调度器接口
type SchedulingPool interface {
	SubmitImage(*ImageInfo)
	SubmitColor(*ImgColor)
}

func NewConcurrent(workCount int) *Concurrent {
	once.Do(func() {
		Scheduler = &Concurrent{
			workerCount: workCount,
			Item: Item{
				inColorChan:   make(chan *ImgColor),
				imageListChan: make(chan *ImageInfo, ImageChanSize),
				downloadChan:  make(chan *ImageInfo),
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
			//case image := <-c.inImageChan:
			//	//下载图片
			//	//上传七牛
			//	//保存数据库
			//	fmt.Println("go image:", i, ">>", image)
			case color := <-c.inColorChan:
				fmt.Println("go color:", i, ">>", color)
			case image := <-c.downloadChan:
				fmt.Println("image:", i, image)
				qiniu.UploadFile()
			case <-c.Ctx.Done():
				fmt.Println("Worker", i, "终止请求.....")
				return
			}
		}
	}()
}

func (c *Concurrent) SubmitImage(image *ImageInfo) {
	c.imageListChan <- image
}

func (c *Concurrent) SubmitColor(color *ImgColor) {
	c.inColorChan <- color
}
