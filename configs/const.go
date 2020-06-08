/*
@Time : 2020/6/8 15:08
@Author : zxr
@File : const
@Software: GoLand
*/
package configs

// 软件信息。
const (
	VERSION   string = "v1.0.0"                                      // 软件版本号
	AUTHOR    string = "zxr"                                         // 软件作者
	NAME      string = "pixabay.com图片抓取"                             // 软件名
	FULL_NAME string = NAME + "_" + VERSION + " （by " + AUTHOR + "）" // 软件全称
)

const (
	ImageType = "image"
	VideoType = "video"
)
