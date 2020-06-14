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
	ApiKey        string
	Id            string
	Page          int    //页码
	Limit         int    //每页的结果数
	Color         string //按颜色属性过滤图像。逗号分隔的值列表可用于选择多个属性 Accepted values: "grayscale", "transparent", "red", "orange", "yellow", "green", "turquoise", "blue", "lilac", "pink", "white", "gray", "black", "brown"
	Type          string //image:图片, video:视频
	Q             string //URL编码的搜索词 示例：“黄色​​+花朵”
	Lang          string //语言代码。zh:中文  默认值:en
	EditorsChoice bool   //选择已获得编辑选择奖的图像,[小编推荐的图片]接受的值true，false
	Order         string //结果如何排序,接受的值:popular,latest。latest:最新
	Category      string //分类 Accepted values: backgrounds, fashion, nature, science, education, feelings, health, people, religion, places, animals, industry, computer, food, sports, transportation, travel, buildings, business, music
	Orientation   string //图像宽于高还是宽于高,接受的值：“所有”，“水平”，“垂直”  Accepted values: "all", "horizontal", "vertical"
	VideoType     string //视频类型过滤结果，接受值:“全部”、“电影”、“动画” Accepted values: "all", "film", "animation"
	ImageType     string //图像类型过滤结果。 Accepted values: "all", "photo", "illustration", "vector"，Default: "all"
}

type HttpParams struct {
	Ua string
}

type Api interface {
	RequestImage()
	RequestVideo()
}

//API请求图片接口返回信息
type ApiImageResp struct {
	Total     int         `json:"total"`
	TotalHits int         `json:"totalHits"`
	Hits      []ItemImage `json:"hits"`
}

//图片解析信息item
type ItemImage struct {
	ID              int    `json:"id"`
	PageURL         string `json:"pageURL"`
	Type            string `json:"type"`
	Tags            string `json:"tags"`
	PreviewURL      string `json:"previewURL"`
	PreviewWidth    int    `json:"previewWidth"`
	PreviewHeight   int    `json:"previewHeight"`
	WebformatURL    string `json:"webformatURL"`
	WebformatWidth  int    `json:"webformatWidth"`
	WebformatHeight int    `json:"webformatHeight"`
	LargeImageURL   string `json:"largeImageURL"`
	ImageWidth      int    `json:"imageWidth"`
	ImageHeight     int    `json:"imageHeight"`
	ImageSize       int    `json:"imageSize"`
	ItemImageStat
	ItemAuthor
}

//作者信息
type ItemAuthor struct {
	UserID       int    `json:"user_id"`
	User         string `json:"user"`
	UserImageURL string `json:"userImageURL"`
}

type ItemImageStat struct {
	Views     int `json:"views"`
	Downloads int `json:"downloads"`
	Favorites int `json:"favorites"`
	Likes     int `json:"likes"`
	Comments  int `json:"comments"`
}

//视频信息 item
type ItemVideo struct {
}
