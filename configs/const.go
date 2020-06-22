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
	ImageType    = "image"
	VideoType    = "video"
	All          = "all"
	OrderPopular = "popular"
	OrderLatest  = "latest"
)


//七牛配置
const (
	QINIU_BKT = "--"
	QINIU_AK  = "--"
	QINIU_SK  = "--"
)


//数据库配置
const (
	DbHost     = "--"
	DbUser     = "--"
	DbPassWd   = "--"
	DbPort     = 
	DbDataBase = "--"
)

//ES配置
const (
	ES_HOST  = "http"
	ES_INDEX = "--"
)


const (
	ApiKey = "hello -- world"
)
