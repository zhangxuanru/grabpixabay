/*
@Time : 2020/6/2 14:40
@Author : zxr
@File : search
@Software: GoLand
*/
package spider

import (
	"errors"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PixSearch struct {
	Html  *string
	Url   string
	Color string
	Dom   *goquery.Document
}

//解析入口搜索页的HTML，
//https://pixabay.com/zh/images/search/?colors=red
func (p *PixSearch) HtmlParser() (err error) {
	if *p.Html == "" {
		return errors.New("body 为空")
	}
	body := *p.Html
	if p.Dom, err = goquery.NewDocumentFromReader(strings.NewReader(body)); err != nil {
		return err
	}
	//总数量
	numText := p.Dom.Find("div.media_list").Find("div>h1").Text()
	numText = strings.TrimSpace(strings.TrimRight(numText, "免费图片"))
	fmt.Println(numText)

	p.ParseHtmlImages()

	//fmt.Printf("HTML:\n\n\n")
	//fmt.Println(*reqRet.Html)

	//logrus.Infoln("开始抓取:", reqRet.Url)

	//解析HTML，发送gorotine请求
	//fmt.Printf("%+v\n\n", reqRet)
	//fmt.Println("Html:", *reqRet.Html)
	return nil
}

//解析出图片地址与链接地址
func (p *PixSearch) ParseHtmlImages() {
	p.Dom.Find("div.search_results").Find("div.item").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(i, ">>>", selection.Text())
		//todo 明天继续，获取图片地址
	})
}
