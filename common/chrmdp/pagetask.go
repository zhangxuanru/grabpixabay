/*
@Time : 2020/6/2 14:45
@Author : zxr
@File : pagetask
@Software: GoLand
*/
package chrmdp

import (
	"io/ioutil"

	"github.com/chromedp/chromedp"
	"github.com/sirupsen/logrus"
)

//请求search 页地址
//https://pixabay.com/zh/images/search/?colors=green
func (r *ReqResult) RequestSearchPage() (err error) {
	logrus.Infoln("开始抓取:", r.Url)
	if isTest, err := r.callTestPage(); isTest == true {
		return err
	}
	err = r.RequestUrl(func(req *ReqResult) chromedp.Tasks {
		return chromedp.Tasks{
			chromedp.Navigate(req.Url),
			// 等待右下角图标加载完成
			chromedp.WaitVisible(`#toTop`, chromedp.ByQuery),
			chromedp.OuterHTML(`body`, req.Html, chromedp.ByQuery),
		}
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"url": r.Url,
		}).Error(err)
		logrus.Infoln("抓取失败,失败原因:", err)
	}
	return err
}

//请求测试文件
func (r *ReqResult) callTestPage() (bool, error) {
	if r.PageType == PageTypeAll {
		byte, err := ioutil.ReadFile("test/html/search.html")
		html := string(byte)
		r.Html = &html
		return true, err
	}
	return false, nil
}
