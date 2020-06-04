/*
@Time : 2020/6/2 14:40
@Author : zxr
@File : search
@Software: GoLand
*/
package spider

import (
	"errors"
	"grabpixabay/app/scheduler"
	"grabpixabay/config"
	"math/rand"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//解析入口搜索页的HTML，
//https://pixabay.com/zh/images/search/?colors=red
func (p *PixSearch) HtmlParser() (err error, nextPage int) {
	if *p.Html == "" {
		return errors.New("body 为空"), 0
	}
	body := *p.Html
	if p.Dom, err = goquery.NewDocumentFromReader(strings.NewReader(body)); err != nil {
		return err, 0
	}
	p.ParseImagesCount()
	p.ParseHtmlImages()
	nextPage = p.ParseHtmlPage()
	return nil, nextPage
}

//解析出图片信息
func (p *PixSearch) ParseHtmlImages() {
	p.Dom.Find("div.search_results").Find("div.item").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		var (
			exists   bool
			srcSet   string
			firstImg *goquery.Selection
		)
		imgInfo := &scheduler.ImageInfo{
			Color:    p.Color,
			ImageSet: make(map[string]string),
		}
		if imgInfo.LinkUrl, exists = selection.Find("a").Eq(0).Attr("href"); exists == false {
			return false
		}
		firstImg = selection.Find("img").Eq(0)
		imgInfo.Alt, _ = firstImg.Attr("alt")
		if imgInfo.ImgSrc, exists = firstImg.Attr("data-lazy"); exists == false {
			imgInfo.ImgSrc, _ = firstImg.Attr("src")
		}
		if srcSet, exists = firstImg.Attr("srcset"); exists == false {
			srcSet, _ = firstImg.Attr("data-lazy-srcset")
		}
		if srcSet != "" {
			srcSetList := strings.Split(strings.TrimSpace(srcSet), ",")
			for _, v := range srcSetList {
				imgSet := strings.Split(strings.TrimSpace(v), " ")
				imgInfo.ImageSet[imgSet[1]] = imgSet[0]
			}
		}
		likeNumText := selection.Find("em").Eq(0).Text()
		favNumText := selection.Find("em").Eq(1).Text()
		comNumText := selection.Find("em").Eq(2).Text()
		if likeNumText != "" {
			imgInfo.LikeNum, _ = strconv.Atoi(strings.TrimSpace(likeNumText))
		}
		if favNumText != "" {
			imgInfo.FavoriteNum, _ = strconv.Atoi(strings.TrimSpace(favNumText))
		}
		if comNumText != "" {
			imgInfo.CommentsNum, _ = strconv.Atoi(strings.TrimSpace(comNumText))
		}
		//将图片信息发送到scheduler
		p.Scheduler.SubmitImage(imgInfo)
		return true
	})
}

//根据颜色 获取图片总数 总数量
func (p *PixSearch) ParseImagesCount() {
	numText := p.Dom.Find("div.media_list").Find("div>h1").Text()
	numText = strings.TrimSpace(strings.TrimRight(numText, "免费图片"))
	if numText != "" {
		count, _ := strconv.Atoi(numText)
		color := &scheduler.ImgColor{
			Color: p.Color,
			Count: count,
		}
		p.Scheduler.SubmitColor(color)
	}
}

//解析是否有下一页,如果有下一页,则返回下一页的页数
func (p *PixSearch) ParseHtmlPage() int {
	var (
		countPage int //总页数
		currPage  int //当前页数
	)
	countPageText := p.Dom.Find("div.paginator > form").First().Text()
	countPageText = strings.TrimSpace(countPageText)
	countPageText = strings.TrimFunc(countPageText, func(r rune) bool {
		if r == 32 || r == 10 || r == 47 {
			return true
		}
		return false
	})
	currPageText, _ := p.Dom.Find("div.paginator > form >input").First().Attr("value")
	//color, _ := p.Dom.Find("div.paginator > form >input").Eq(1).Attr("value")
	currPageText = strings.TrimSpace(currPageText)
	if countPageText == "" || currPageText == "" {
		return -1
	}
	if len(currPageText) > 0 {
		currPage, _ = strconv.Atoi(currPageText)
	}
	if len(countPageText) > 0 {
		countPage, _ = strconv.Atoi(countPageText)
	}
	if currPage >= countPage {
		return 0
	}
	//测试环境下，临时返回页数
	if config.SYSLEVEL == config.DEBUG_LEVEL && currPage < countPage {
		nextPage := (currPage + 1 + rand.Intn(10)) * 5
		if nextPage > 25 {
			return 0
		}
		return nextPage
	}
	if currPage < countPage {
		return currPage + 1
	}
	return 0
}
