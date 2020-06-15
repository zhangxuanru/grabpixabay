/*
@Time : 2020/5/29 17:36
@Author : zxr
@File : conf
@Software: GoLand
*/
package configs

var GConf *Config

type Config struct {
	Colors            []string //所有颜色列表
	ImageType         []string //图片类型
	Orientation       []string //图像宽高
	Category          []string //所有分类列表
	WorkerCount       int      //同时开几个worker处理抓取信息
	ItemQueueMaxLimit int      //队列存储的最大量
}

func AppConfig() *Config {
	GConf = &Config{
		Colors:            []string{"grayscale", "transparent", "red", "orange", "yellow", "green", "turquoise", "blue", "lilac", "pink", "white", "gray", "black", "brown"},
		ImageType:         []string{"photo", "illustration", "vector"},
		Orientation:       []string{"horizontal", "vertical"},
		Category:          []string{"backgrounds", "fashion", "nature", "science", "education", "feelings", "health", "people", "religion", "places", "animals", "industry", "computer", "food", "sports", "transportation", "travel", "buildings", "business", "music"},
		WorkerCount:       100,
		ItemQueueMaxLimit: 50000,
	}
	return GConf
}
