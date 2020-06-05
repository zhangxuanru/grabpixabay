/*
@Time : 2020/6/5 15:04
@Author : zxr
@File : types
@Software: GoLand
*/
package storage

type ImgColor struct {
	Color      string //颜色
	Count      int    //源站上的图片总数
	SuccessNum int    //下载成功的数
	FailNum    int    //下载失败的数
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
