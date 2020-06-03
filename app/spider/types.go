/*
@Time : 2020/6/3 16:33
@Author : zxr
@File : types
@Software: GoLand
*/
package spider

import "github.com/PuerkitoBio/goquery"

//首页HTML结构体
type PixSearch struct {
	Html  *string
	Url   string
	Color string
	Dom   *goquery.Document
}

//图片结构体
type ImageInfo struct {
	LinkUrl     string            //图片链接地址
	Alt         string            //图片文字提示
	Tags        []ImageTag        //图片标签
	ImgSrc      string            //图片地址
	ImageSet    map[string]string //img-set图片集合
	LikeNum     int               //喜欢数
	FavoriteNum int               //收藏数
	CommentsNum int               //评论数
	Color       string            //颜色
}

type ImageTag struct {
	Title   string //标签标题
	LinkUrl string //标签链接地址
}
