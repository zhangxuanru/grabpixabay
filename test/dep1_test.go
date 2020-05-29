/*
@Time : 2020/5/26 17:04
@Author : zxr
@File : dep1_test
@Software: GoLand
*/
package test

import (
	"context"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

// 必应搜索
func TestDep1(t *testing.T) {
	var buf []byte

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		//访问打开必应页面
		chromedp.Navigate(`https://cn.bing.com/?mkt=zh-CN`),
		// 等待右下角图标加载完成
		chromedp.WaitVisible(`#sh_cp_in`),
		//搜索框内输入zhangguanzhang
		chromedp.SendKeys(`#sb_form_q`, `周星驰`, chromedp.ByID),
		// 点击搜索图标
		chromedp.Click(`#sb_form_go`, chromedp.NodeVisible),
		// 获取第一个搜索结构的超链接
		//document.querySelector("#b_results > li:nth-child(2) > h2 > a")
		//document.querySelector("#b_results > li:nth-child(1)")
		chromedp.OuterHTML(`#b_results`, &example, chromedp.ByQuery),
		chromedp.CaptureScreenshot(&buf),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("fullScreenshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("example: %s", example)
}
