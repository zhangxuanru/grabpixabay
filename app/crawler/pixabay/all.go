/*
@Time : 2020/6/1 16:57
@Author : zxr
@File : imgsearch
@Software: GoLand
*/
package pixabay

import (
	"fmt"
	"grabpixabay/app/scheduler"
	"grabpixabay/app/spider"
	"grabpixabay/common/chrmdp"
	"grabpixabay/config"

	"github.com/sirupsen/logrus"
)

//执行全站图片所有抓取
//https://pixabay.com/zh/images/search/
type CrawlerAll struct {
	Title      string
	PixRequest *PixRequest
	scheduler  *scheduler.SchedulPool
}

func NewCrawlerAll(req *PixRequest) *CrawlerAll {
	return &CrawlerAll{
		Title:      "全站抓取",
		PixRequest: req,
	}
}

//入口
func (c *CrawlerAll) Start() {
	c.StartWorker()
	for _, color := range config.GConf.Colors {
		select {
		case <-c.PixRequest.Cxt.Done():
			fmt.Println("终止请求.....")
			return
		default:
			_ = c.CrawlerColorPage(color)
		}
	}
}

//抓取所有， 根据颜色发起请求
//https://pixabay.com/zh/images/search/?colors=green
func (c *CrawlerAll) CrawlerColorPage(color string) (err error) {
	url := c.PixRequest.HostUrl + "?colors=" + color
	reqRet := chrmdp.NewReqResult(url, chrmdp.PageTypeAll)
	if err = reqRet.RequestSearchPage(); err != nil {
		return
	}
	query := &spider.PixSearch{
		Html:      reqRet.Html,
		Url:       reqRet.Url,
		Color:     color,
		Ctx:       c.PixRequest.Cxt,
		Can:       c.PixRequest.Can,
		Scheduler: c.scheduler,
	}
	if err = query.HtmlParser(); err != nil {
		logrus.Error(err)
		return
	}
	logrus.Infoln("抓取结束:", reqRet.Url)
	return nil
}

//启动worker
func (c *CrawlerAll) StartWorker() {
	worker := scheduler.NewConcurrent(config.GConf.WorkerCount)
	worker.Ctx = c.PixRequest.Cxt
	worker.Cancel = c.PixRequest.Can
	worker.Run()
	c.scheduler = worker
}
