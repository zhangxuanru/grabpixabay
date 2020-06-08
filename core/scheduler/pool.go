/*
@Time : 2020/6/8 15:43
@Author : zxr
@File : pool
@Software: GoLand
*/
package scheduler

import "fmt"

func NewConcurrent(workerCount int) *Concurrent {
	return &Concurrent{
		itemImageChan: make(chan *ItemImage),
		itemVideoChan: make(chan *ItemVideo),
		workerCount:   workerCount,
	}
}

func (c *Concurrent) Run() {
	for i := 0; i < c.workerCount; i++ {
		c.createWorker(i)
	}
}

func (c *Concurrent) createWorker(i int) {
	go func() {
		for {
			select {
			case image := <-c.itemImageChan:
				fmt.Printf("go..%d, rev:%+v\n", i, image)
			}
		}
	}()
}

func (c *Concurrent) SubmitImageItem(item *ItemImage) {
	c.itemImageChan <- item
}
