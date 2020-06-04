/*
@Time : 2020/6/1 11:58
@Author : zxr
@File : crawler
@Software: GoLand
*/
package pixabay

import (
	"fmt"
	"grabpixabay/app/scheduler"
	"grabpixabay/config"
)

func NewPixRequest() *PixRequest {
	return &PixRequest{}
}

//判断抓取的类型
func (p *PixRequest) CrawPixType(crawType string) {
	switch crawType {
	case config.TYPE_ALL:
		p.RunAll()
	case config.TYPE_LATEST:
		p.RunLatest()
	case config.TYPE_SIFT:
		p.RunSift()
	case config.TYPE_PIC:
		p.RunPic()
	default:
		fmt.Println("type is Undefined")
		return
	}
}

//执行全站图片所有抓取
func (p *PixRequest) RunAll() {
	p.StartWorker()
	NewCrawlerAll(p).Start()
}

func (p *PixRequest) RunLatest() {
	fmt.Println("latest 待开发....")
}

func (p *PixRequest) RunSift() {
	fmt.Println("sift 待开发....")
}

func (p *PixRequest) RunPic() {
	fmt.Println("pic 待开发....")
}

//启动调度器
func (p *PixRequest) StartWorker() {
	schedule := scheduler.NewConcurrent(config.GConf.WorkerCount)
	schedule.Ctx = p.Cxt
	schedule.Cancel = p.Can
	schedule.Run()
	p.SchPool = schedule
}
