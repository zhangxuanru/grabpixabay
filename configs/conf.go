/*
@Time : 2020/5/29 17:36
@Author : zxr
@File : conf
@Software: GoLand
*/
package configs

var GConf *Config

type Config struct {
	Colors      []string
	WorkerCount int //同时开几个worker处理抓取信息

}

func AppConfig() *Config {
	GConf = &Config{
		Colors:      []string{"grayscale", "transparent", "red", "orange", "yellow", "green", "turquoise", "blue", "lilac", "pink", "white", "gray", "black", "brown"},
		WorkerCount: 100,
	}
	return GConf
}
