/*
@Time : 2020/5/26 17:45
@Author : zxr
@File : dep2_test
@Software: GoLand
*/
package test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

//百度，输入搜索词搜索
func TestDep2(t *testing.T) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)
	ctx := context.Background()
	c, cc := chromedp.NewExecAllocator(ctx, options...)
	defer cc()

	ctx, cancel := chromedp.NewContext(c,
		chromedp.WithLogf(log.Printf))
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	// navigate to a page, wait for an element, click
	var html string
	var key = "科比"

	err := chromedp.Run(ctx, getBaidu(&html, key))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("html: %s", html)
}

func getBaidu(html *string, key string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://www.baidu.com/`),
		// 等待右下角图标加载完成
		chromedp.WaitVisible(`#form`, chromedp.ByQuery),
		//搜索框内输入
		//chromedp.SendKeys(`#kw`, `科比`, chromedp.ByID),
		chromedp.SetValue(`#kw`, key, chromedp.ByID),

		// 点击搜索图标
		chromedp.Submit(`#su`, chromedp.ByID),
		chromedp.Sleep(1 * time.Second),
		chromedp.OuterHTML(`body`, html, chromedp.ByQuery),
	}
}
