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
	PIX_HOST    = "pixabay" //host名称
	TYPE_ALL    = "all"     //全站抓取
	TYPE_LATEST = "latest"  //获取最新
	TYPE_SIFT   = "sift"    //获取小编精选
	TYPE_PIC    = "pic"     //获取图片
)

const (
	DEBUG_LEVEL   = "DEBUG"
	PRODUCT_LEVEL = "PRODUCT"
	SYSLEVEL      = DEBUG_LEVEL
)
