/*
@Time : 2020/6/8 17:39
@Author : zxr
@File : request
@Software: GoLand
*/
package api

import (
	"bytes"
	"fmt"
)

//请求图片信息的API
func NewImageRequest() *ImageRequest {
	return &ImageRequest{}
}

//请求APi
func (r *ImageRequest) Request() {
	r.BuildApiUrl()
}

//拼接URL参数
func (r *ImageRequest) BuildParams() string {
	var buf bytes.Buffer
	if r.Size > 0 {
		buf.WriteString(fmt.Sprintf("&per_page=%d", r.Size))
	}
	if r.Page > 0 {
		buf.WriteString(fmt.Sprintf("&page=%d", r.Page))
	}
	if r.Color != "" {
		buf.WriteString(fmt.Sprintf("&colors=%s", r.Color))
	}
	if r.Q != "" {
		buf.WriteString(fmt.Sprintf("&q=%s", r.Q))
	}
	if r.Category != "" {
		buf.WriteString(fmt.Sprintf("&category=%s", r.Category))
	}
	if r.Order != "" {
		buf.WriteString(fmt.Sprintf("&order=%s", r.Order))
	}
	if r.EditorsChoice == true {
		buf.WriteString(fmt.Sprintf("&editors_choice=%t", r.EditorsChoice))
	}
	if r.ImageType != "" {
		buf.WriteString(fmt.Sprintf("&image_type=%s", r.ImageType))
	}
	return buf.String()
}

//获取API URL
func (r *ImageRequest) BuildApiUrl() {
	params := r.BuildParams()
	url := r.ApiUrl + params
	fmt.Println(url)
}
