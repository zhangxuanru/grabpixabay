/*
@Time : 2020/5/29 17:36
@Author : zxr
@File : conf
@Software: GoLand
*/
package config

var GConf *Config

type Config struct {
	HostMap map[string]string
	TypeMap map[string]string
	Colors  []string
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
		Colors: []string{"red", "orange", "yellow", "green", "turquoise", "blue", "lilac", "pink", "white", "gray", "black", "brown"},
	}
	return GConf
}
