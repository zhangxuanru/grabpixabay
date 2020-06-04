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
	"sync"
)

var Scheduler *Concurrent
var once sync.Once

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
				inImageChan: make(chan *ImageInfo),
				inColorChan: make(chan *ImgColor),
			},
		}
	})
	return Scheduler
}

func (c *Concurrent) Run() {
	if c.WorkActive == true {
		return
	}
	for i := 0; i < c.workerCount; i++ {
		c.createWorker(i)
	}
	c.WorkActive = true
}

func (c *Concurrent) createWorker(i int) {
	go func() {
		for {
			select {
			case image := <-c.inImageChan:
				fmt.Println("go image:", i, ">>", image)
			case color := <-c.inColorChan:
				fmt.Println("go color:", i, ">>", color)
			case <-c.Ctx.Done():
				fmt.Println("Worker", i, "终止请求.....")
				return
			}
		}
	}()
}

func (c *Concurrent) SubmitImage(image *ImageInfo) {
	c.inImageChan <- image
}

func (c *Concurrent) SubmitColor(color *ImgColor) {
	c.inColorChan <- color
}
