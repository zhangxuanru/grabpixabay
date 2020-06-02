/*
@Time : 2020/6/1 16:57
@Author : zxr
@File : imgsearch
@Software: GoLand
*/
package pixabay

import (
	"grabpixabay/app/spider"
	"grabpixabay/common/chrmdp"
	"grabpixabay/config"
	"time"

	"github.com/sirupsen/logrus"
)

//执行全站图片所有抓取
//https://pixabay.com/zh/images/search/
type CrawlerAll struct {
	Title      string
	PixRequest *PixRequest
}

func NewCrawlerAll(req *PixRequest) *CrawlerAll {
	return &CrawlerAll{
		Title:      "全站抓取",
		PixRequest: req,
	}
}

//抓取所有， 根据颜色发起请求
//https://pixabay.com/zh/images/search/?colors=green
func (c *CrawlerAll) CrawlerAll() {
	for _, color := range config.GConf.Colors {
		url := c.PixRequest.HostUrl + "?colors=" + color
		reqRet := chrmdp.NewReqResult(url, chrmdp.PageTypeAll)
		if err := reqRet.RequestSearchPage(); err != nil {
			continue
		}
		query := &spider.PixSearch{
			Html:  reqRet.Html,
			Url:   reqRet.Url,
			Color: color,
		}
		if err := query.HtmlParser(); err != nil {
			logrus.Error(err)
		}
		logrus.Infoln("抓取结束:", reqRet.Url)
		time.Sleep(1 * time.Second)
		return
	}
}
