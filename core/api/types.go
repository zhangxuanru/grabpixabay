/*
@Time : 2020/6/8 17:22
@Author : zxr
@File : types
@Software: GoLand
*/
package api

//https://pixabay.com/api/docs/

type RequestInfo struct {
	ApiUrl        string
	Id            string
	Page          int    //页码
	Size          int    //每页的结果数
	Color         string //按颜色属性过滤图像。逗号分隔的值列表可用于选择多个属性 Accepted values: "grayscale", "transparent", "red", "orange", "yellow", "green", "turquoise", "blue", "lilac", "pink", "white", "gray", "black", "brown"
	Type          string //image:图片, video:视频
	Q             string //URL编码的搜索词 示例：“黄色​​+花朵”
	Lang          string //语言代码。zh:中文  默认值:en
	EditorsChoice bool   //选择已获得编辑选择奖的图像,[小编推荐的图片]接受的值true，false
	Order         string //结果如何排序,接受的值:popular,latest。latest:最新
	Category      string //分类 Accepted values: backgrounds, fashion, nature, science, education, feelings, health, people, religion, places, animals, industry, computer, food, sports, transportation, travel, buildings, business, music
}

//图片资料请求信息
type ImageRequest struct {
	ImageType string //图像类型过滤结果。 Accepted values: "all", "photo", "illustration", "vector"，Default: "all"
	RequestInfo
}

//视频资料请求信息，todo 下期做
type VideoRequest struct {
	VideoType string //视频类型过滤结果，接受值:“全部”、“电影”、“动画” Accepted values: "all", "film", "animation"
	RequestInfo
}

type Api interface {
	Request()
}