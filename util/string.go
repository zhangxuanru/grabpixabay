/*
@Time : 2020/6/8 16:57
@Author : zxr
@File : bytes
@Software: GoLand
*/
package util

import "strings"

//判断字符串是否在字符串数组中
func InStrings(arr []string, val string) bool {
	for _, v := range arr {
		if strings.EqualFold(v, val) {
			return true
		}
	}
	return false
}
