/*
@Time : 2020/6/3 17:24
@Author : zxr
@File : worker
@Software: GoLand
*/
package scheduler

import "context"

type SchedulPool interface {
	SubmitImage(*ImageInfo)
	SubmitColor(*ImgColor)
	Run()
}

type Work struct {
	inImageChan  chan *ImageInfo //
	outImageChan chan ImageInfo
	inColorChan  chan *ImgColor
	outColorChan chan ImgColor
}

type Concurrent struct {
	workerCount int
	WorkActive  bool
	Ctx         context.Context
	Cancel      context.CancelFunc
	Work
}

type ImgColor struct {
	Color      string //颜色
	Count      int    //源站上的图片总数
	SuccessNum int    //下载成功的数
	FailNum    int    //下载失败的数
}
