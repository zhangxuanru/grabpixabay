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

	"github.com/chromedp/chromedp"
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
func (c *CrawlerAll) CrawlerAll() {
	var err error
	//https://pixabay.com/zh/images/search/?colors=green
	for _, color := range config.GConf.Colors {
		//todo 明天继续
	}
	reqRet := chrmdp.NewReqResult(c.PixRequest.HostUrl)
	err = reqRet.RequestUrl(func(req *chrmdp.ReqResult) chromedp.Tasks {
		return chromedp.Tasks{
			chromedp.Navigate(req.Url),
			// 等待右下角图标加载完成
			chromedp.WaitVisible(`#toTop`, chromedp.ByQuery),
			chromedp.OuterHTML(`body`, req.Html, chromedp.ByQuery),
		}
	})
	if err != nil {
		logrus.Error(err)

	}

	fmt.Printf("%+v\n\n", reqRet)
	fmt.Println("Html:", *reqRet.Html)
}
