/*
@Time : 2020/5/29 16:02
@Author : zxr
@File : config
@Software: GoLand
*/
package config

// 软件信息。
const (
	VERSION   string = "v1.0.0"                                      // 软件版本号
	AUTHOR    string = "zxr"                                         // 软件作者
	NAME      string = "pixabay.com图片抓取"                             // 软件名
	FULL_NAME string = NAME + "_" + VERSION + " （by " + AUTHOR + "）" // 软件全称
	TAG       string = "zxr_px"                                      // 软件标识符
)

//命令行信息
const (
	PIX_HOST    = "pixabay"
	TYPE_ALL    = "all"
	TYPE_LATEST = "latest"
	TYPE_SIFT   = "sift"
	TYPE_PIC    = "pic"
)
