/*
@Time : 2020/6/2 14:45
@Author : zxr
@File : pagetask
@Software: GoLand
*/
package chrmdp

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/sirupsen/logrus"
)

//请求search 页地址
//https://pixabay.com/zh/images/search/?colors=green
func (r *ReqResult) RequestSearchPage() (err error) {
	logrus.Infoln("开始抓取:", r.Url)
	if isTest, err := r.TestHtmlFile(); isTest == true {
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

//请求图片详情页地址
//https://pixabay.com/zh/photos/strawberries-ripe-heart-by-heart-5210753/
func (r *ReqResult) RequestPicDetailPage() (err error) {
	logrus.Infoln("开始抓取图片详情页:", r.Url)
	if isTest, err := r.TestHtmlFile(); isTest == true {
		return err
	}
	err = r.RequestUrl(func(req *ReqResult) chromedp.Tasks {
		return chromedp.Tasks{
			chromedp.Navigate(req.Url),
			//cookie
			SetCookie("DSID", "AAO-7r6vjf720W6fq7QkDFfXQNqbH5YKJQFeLVUHu7tkwZpH3q3_BEUf3HJKtEvi1NtoKaNJx73Q4bXrMsNpc2LdU0oAPhrouXsZiTVPmiFHZRJlxZvDUj8", ".doubleclick.net", "/", false, false),
			SetCookie("__cfduid", "da916b950e1b7750a11fb0c76b4a90c891591348720", ".pixabay.com", "/", false, false),
			SetCookie("__cfduid", "d1a0b9978ff7979c74b27607ceb87380f1591346493", ".appboycdn.com", "/", false, false),
			SetCookie("_ga", "GA1.2.252412743.1591349371", ".pixabay.com", "/", false, false),
			SetCookie("_gat_UA-20223345-1", "1", ".pixabay.com", "/", false, false),
			SetCookie("_gid", "GA1.2.1975057857.1591349371", ".pixabay.com", "/", false, false),
			SetCookie("ab.storage.sessionId.a5fd3939-90ba-4678-86ee-0a0c8fb3d061", "%7B%22g%22%3A%2262026e9d-3627-9a03-414d-8c2a9f94afd1%22%2C%22e%22%3A1591351240068%2C%22c%22%3A1591347144381%2C%22l%22%3A1591349440068%7D", ".pixabay.com", "/", false, false),
			SetCookie("anonymous_user_id", "45d31f2b-21e9-4092-beca-46cd9eaeef90", "pixabay.com", "/", false, false),
			SetCookie("bsaw", "7493499201780585951", "srv.carbonads.net", "/", false, false),
			SetCookie("client_width", "654", "pixabay.com", "/", false, false),
			SetCookie("dwf_attribution_template_ads", "True", "pixabay.com", "/", false, false),
			SetCookie("id", "22aaeb59ebad00ac||t=1545985344|et=730|cs=002213fd4848262d08275dbe14", ".doubleclick.net", "/", false, false),
			SetCookie("is_human", "1", "pixabay.com", "/", false, false),
			SetCookie("lang", "zh", "pixabay.com", "/", false, false),

			// 等待右下角图标加载完成
			chromedp.WaitVisible(`#toTop`, chromedp.ByQuery),
			chromedp.OuterHTML(`body`, req.Html, chromedp.ByQuery),
		}
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"url": r.Url,
		}).Error(err)
		logrus.Infoln("图片详情页抓取失败,失败原因:", err)
	}
	return err
}

func SetCookie(name, value, domain, path string, httpOnly, secure bool) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
		success, err := network.SetCookie(name, value).
			WithExpires(&expr).
			WithDomain(domain).
			WithPath(path).
			WithHTTPOnly(httpOnly).
			WithSecure(secure).
			Do(ctx)
		if err != nil {
			return err
		}
		if !success {
			return fmt.Errorf("could not set cookie %s", name)
		}
		return nil
	})
}

//请求测试文件
func (r *ReqResult) TestHtmlFile() (bool, error) {
	if r.TestFile == false {
		return false, nil
	}
	if r.PageType == PageTypeAll {
		content, err := ioutil.ReadFile("test/html/search.html")
		html := string(content)
		r.Html = &html
		return true, err
	}
	return false, nil
}
