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
	"sync"

	"github.com/sirupsen/logrus"
)

func NewConcurrent(workerCount int, ctx context.Context, can context.CancelFunc) *Concurrent {
	return &Concurrent{
		itemImageChan: make(chan ItemImage),
		itemVideoChan: make(chan *ItemVideo),
		workerCount:   workerCount,
		Ctx:           ctx,
		Can:           can,
		Wg:            &sync.WaitGroup{},
	}
}

func (c *Concurrent) Run() {
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
				logrus.Infof("go worker %d, rev:%+v\n", i, image)
				c.Wg.Done()
			case <-c.Ctx.Done():
				fmt.Println("Worker", i, "终止请求.....")
				return
			}
		}
	}()
}

func (c *Concurrent) SubmitImageItem(item ItemImage) {
	c.itemImageChan <- item
}

func (c *Concurrent) Wait() {
	c.Wg.Wait()
}

func (c *Concurrent) AddWgNum(n int) {
	c.Wg.Add(n)
}
