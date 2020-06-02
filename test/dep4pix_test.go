/*
@Time : 2020/5/26 17:45
@Author : zxr
@File : dep2_test
@Software: GoLand
*/
package test

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"testing"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"

	"github.com/chromedp/chromedp"
)

//pixabay
func TestPixabay(t *testing.T) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
		chromedp.Flag("headless", false),
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

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 40*time.Second)
	defer cancel()
	// navigate to a page, wait for an element, click
	var html string

	var buf []byte
	err := chromedp.Run(ctx, getPixabay(&html, &buf, 90))
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("pix.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("html: %s", html)
}

func getPixabay(html *string, res *[]byte, quality int64) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://pixabay.com/zh/images/search/?colors=red`),
		// 等待右下角图标加载完成
		//chromedp.WaitVisible(`#form`, chromedp.ByQuery),
		//搜索框内输入
		//chromedp.SendKeys(`#kw`, `科比`, chromedp.ByID),
		//chromedp.SetValue(`#kw`, key, chromedp.ByID),

		// 点击搜索图标
		//chromedp.Submit(`#su`, chromedp.ByID),
		chromedp.Sleep(1 * time.Second),
		chromedp.OuterHTML(`body`, html, chromedp.ByQuery),
		SetCookie("ab.storage.sessionId.a5fd3939-90ba-4678-86ee-0a0c8fb3d061", `{"v":{"g":"334e74f5-83bf-8fad-cae2-8eedeee75e92","e":1591086527331,"c":1591083240487,"l":1591084727331}}`, "pixabay.com", "/", false, false),

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
