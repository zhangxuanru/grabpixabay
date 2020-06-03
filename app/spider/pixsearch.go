/*
@Time : 2020/6/2 14:40
@Author : zxr
@File : search
@Software: GoLand
*/
package spider

import (
	"errors"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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
	p.ParseHtmlImages()
	return nil
}

//解析出图片信息
func (p *PixSearch) ParseHtmlImages() {
	p.Dom.Find("div.search_results").Find("div.item").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		var (
			exists   bool
			srcSet   string
			firstImg *goquery.Selection
		)
		imgInfo := &ImageInfo{
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
		//fmt.Printf("%+v\n\n", imgInfo)
		return true
	})
}
