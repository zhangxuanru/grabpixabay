/*
@Time : 2020/6/9 16:34
@Author : zxr
@File : format
@Software: GoLand
*/
package scheduler

import (
	jsoniter "github.com/json-iterator/go"
	"grabpixabay/core/api"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//转为ApiImageResp结构
func ToApiImageResp(data []byte) (resp *api.ApiImageResp, err error) {
	resp = &api.ApiImageResp{}
	err = json.Unmarshal(data, resp)
	return
}

//将查询参数加到每个图片信息上
func buildReqParams(item api.ItemImage, reqObj *api.RequestInfo) (result api.ItemImage) {
	if reqObj.Color != "" {
		item.Color = reqObj.Color
	}
	if reqObj.Category != "" {
		item.Category = reqObj.Category
	}
	if reqObj.EditorsChoice == true {
		item.EditorsChoice = true
	}
	if reqObj.Orientation != "" {
		item.Orientation = reqObj.Orientation
	}
	return item
}
