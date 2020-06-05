/*
@Time : 2020/6/5 16:20
@Author : zxr
@File : worker
@Software: GoLand
*/
package pixabay

import (
	"fmt"
	"grabpixabay/app/storage"
	"grabpixabay/config"
	"time"
)

type worker struct {
	imageChan     chan *storage.ImageInfo
	imagePageList chan *storage.ImageInfo
	PixRequest    *PixRequest
}

func NewWorker(pix *PixRequest) *worker {
	return &worker{
		imageChan:     make(chan *storage.ImageInfo),
		imagePageList: make(chan *storage.ImageInfo, config.GConf.MaxImageListSize),
		PixRequest:    pix,
	}
}

func (w *worker) StartWorker() {
	for i := 0; i < config.GConf.ImageDetailWorker; i++ {
		w.createPageWorker(i)
	}
	go func() {
		for image := range w.imagePageList {
			w.SubmitImage(image)
		}
	}()
}

func (w *worker) SubmitImage(image *storage.ImageInfo) {
	time.Sleep(5 * time.Second)
	w.imageChan <- image
}

//添加到队列中
func (w *worker) AddImage(image *storage.ImageInfo) {
	w.imagePageList <- image
}

//创建抓取图片详情页的worker
func (w *worker) createPageWorker(i int) {
	go func() {
		for {
			select {
			case image := <-w.imageChan:
				fmt.Println("image detail:", i, ">>", image)
				fmt.Println("count image:", len(w.imagePageList))
				NewCrawlerAll(w.PixRequest).CrawlerImageDetail(image)
			case <-w.PixRequest.Cxt.Done():
				fmt.Println("image Detail Worker", i, "终止请求.....")
				return
			}
		}
	}()
}
