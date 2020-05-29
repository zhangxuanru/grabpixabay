/*
@Time : 2020/5/29 18:25
@Author : zxr
@File : time
@Software: GoLand
*/
package util

import "time"

const APP_LOCAT = "Asia/Shanghai" //	时区

//获取当前unix时间
func GetNowUnixTime() int64 {
	location, _ := time.LoadLocation(APP_LOCAT)
	return time.Now().In(location).Unix()
}
