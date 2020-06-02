/*
@Time : 2020/6/1 16:57
@Author : zxr
@File : imgsearch
@Software: GoLand
*/
package pixabay

import (
	"fmt"
	"grabpixabay/common/chrmdp"
	"grabpixabay/config"

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
		reqRet := chrmdp.NewReqResult(url)
		if err := reqRet.RequestSearchPage(); err != nil {
			continue
		}
		fmt.Printf("%+v\n", reqRet)
		fmt.Printf("HTML:\n\n\n")
		fmt.Println(*reqRet.Html)

		logrus.Infoln("开始抓取:", reqRet.Url)

		//解析HTML，发送gorotine请求
		//fmt.Printf("%+v\n\n", reqRet)
		//fmt.Println("Html:", *reqRet.Html)
		return
	}
}
