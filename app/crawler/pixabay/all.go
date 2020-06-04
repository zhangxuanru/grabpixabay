/*
@Time : 2020/6/1 16:57
@Author : zxr
@File : imgsearch
@Software: GoLand
*/
package pixabay

import (
	"fmt"
	"grabpixabay/app/spider"
	"grabpixabay/common/chrmdp"
	"grabpixabay/config"

	"github.com/sirupsen/logrus"
)

//执行全站图片所有抓取
//https://pixabay.com/zh/images/search/
func NewCrawlerAll(req *PixRequest) *CrawlerAll {
	return &CrawlerAll{
		Title:      "全站抓取",
		PixRequest: req,
		VisitUrl:   make(map[string]struct{}),
	}
}

//入口
func (c *CrawlerAll) Start() {
	for _, color := range config.GConf.Colors {
		select {
		case <-c.PixRequest.Cxt.Done():
			fmt.Println("终止请求.....")
			return
		default:
			_ = c.CrawlerColorPage(color, 1)
		}
	}
}

//抓取所有， 根据颜色发起请求
//https://pixabay.com/zh/images/search/?colors=green
func (c *CrawlerAll) CrawlerColorPage(color string, pag int) (err error) {
	var (
		nextPage int
		url      string
		reqResp  *chrmdp.ReqResult
	)
	url = c.PixRequest.HostUrl + "?colors=" + color
	if pag > 1 {
		url += "&pagi=" + fmt.Sprintf("%d", pag)
	}
	reqResp = chrmdp.NewReqResult(url, chrmdp.PageTypeAll)
	if c.isDuplicate(url) == false {
		if err = reqResp.RequestSearchPage(); err != nil {
			return
		}
	} else {
		logrus.Infoln(reqResp.Url, "重复请求....")
	}
	query := &spider.PixSearch{
		Html:      reqResp.Html,
		Url:       reqResp.Url,
		Color:     color,
		Ctx:       c.PixRequest.Cxt,
		Can:       c.PixRequest.Can,
		Scheduler: c.PixRequest.SchPool,
	}
	if err, nextPage = query.HtmlParser(); err != nil {
		if c.CurrPage > 0 {
			nextPage = c.CurrPage + 1
		}
		logrus.Error(err)
	}
	if nextPage < 1 {
		return
	}
	if c.PixRequest.Page > 0 && nextPage <= c.PixRequest.Page && nextPage > 0 {
		c.CurrPage = nextPage
		c.CrawlerColorPage(query.Color, nextPage)
	}
	if nextPage > 0 && c.PixRequest.Page == 0 {
		c.CurrPage = nextPage
		c.CrawlerColorPage(query.Color, nextPage)
	}
	logrus.Infoln("抓取结束:", query.Url)
	return nil
}

//判断是否访问过
func (c *CrawlerAll) isDuplicate(url string) bool {
	if _, ok := c.VisitUrl[url]; ok {
		return true
	}
	c.VisitUrl[url] = struct{}{}
	return false
}
