/*
@Time : 2020/5/26 17:45
@Author : zxr
@File : dep2_test
@Software: GoLand
*/
package test

import (
	"context"
	"io/ioutil"
	"log"
	"math"
	"testing"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"

	"github.com/chromedp/chromedp"
)

//百度，输入搜索词搜索并截图
func TestDep3(t *testing.T) {
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
	var buf []byte
	err := chromedp.Run(ctx, getBaidu2(&html, key, &buf, 90))
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("baidu.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("html: %s", html)
}

func getBaidu2(html *string, key string, res *[]byte, quality int64) chromedp.Tasks {
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
		chromedp.Sleep(1 * time.Second),

		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}
