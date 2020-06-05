/*
@Time : 2020/5/29 17:36
@Author : zxr
@File : conf
@Software: GoLand
*/
package config

var GConf *Config

type Config struct {
	HostMap           map[string]string
	TypeMap           map[string]string
	Colors            []string
	WorkerCount       int //同时开几个worker处理抓取信息
	ImageDetailWorker int //同时开几个worker处理打开图片详情页抓取信息
	MaxImageListSize  int //图片信息待下载队列，最多存储的个数
}

func AppConfig() *Config {
	GConf = &Config{
		TypeMap: map[string]string{
			TYPE_ALL:    TYPE_ALL,
			TYPE_LATEST: TYPE_LATEST,
			TYPE_PIC:    TYPE_PIC,
			TYPE_SIFT:   TYPE_SIFT,
		},
		HostMap: map[string]string{
			PIX_HOST: "https://pixabay.com/zh/images/search/",
		},
		Colors:            []string{"red", "orange", "yellow", "green", "turquoise", "blue", "lilac", "pink", "white", "gray", "black", "brown"},
		WorkerCount:       100,
		ImageDetailWorker: 5,
		MaxImageListSize:  50000,
	}
	return GConf
}
