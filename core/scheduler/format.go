/*
@Time : 2020/6/9 16:34
@Author : zxr
@File : format
@Software: GoLand
*/
package scheduler

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//转为ApiImageResp结构
func ToApiImageResp(data []byte) (resp *ApiImageResp, err error) {
	resp = &ApiImageResp{}
	err = json.Unmarshal(data, resp)
	return
}
