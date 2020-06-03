/*
@Time : 2020/6/1 18:41
@Author : zxr
@File : request
@Software: GoLand
*/
package chrmdp

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

const UA = `Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`
const timeOut = 30 * time.Second

const (
	PageTypeAll = "all"
)

//请求URL 返回HTML
type ReqResult struct {
	Html     *string
	Url      string
	timeOut  time.Duration
	Ua       string
	Headless bool
	PageType string
	TestFile bool
}

type reqFun func(req *ReqResult) chromedp.Tasks

func NewReqResult(url string, pageType string) *ReqResult {
	var html string
	return &ReqResult{
		Url:      url,
		timeOut:  time.Duration(timeOut),
		Ua:       UA,
		Html:     &html,
		Headless: false,
		PageType: pageType,
		TestFile: true,
	}
}

func (r *ReqResult) RequestUrl(f reqFun) error {
	options := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent(r.Ua),
		chromedp.Flag("headless", r.Headless), //以有头方式运行
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("enable-automation", true),
		chromedp.Flag("disable-extensions", true),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	ctx := context.Background()
	c, cc := chromedp.NewExecAllocator(ctx, options...)
	defer cc()
	ctx, cancel := chromedp.NewContext(c,
		chromedp.WithLogf(log.Printf))
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, r.timeOut)
	defer cancel()
	return chromedp.Run(ctx, f(r))
}
