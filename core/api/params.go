/*
@Time : 2020/6/9 11:12
@Author : zxr
@File : params
@Software: GoLand
*/
package api

import (
	"bytes"
	"fmt"
	"grabpixabay/configs"
)

const ImageApiUrl = "https://pixabay.com/api/"
const VideoApiUrl = "https://pixabay.com/api/videos/"

//api doc : https://pixabay.com/api/docs/
//拼接URL参数
func (r *RequestInfo) buildParams() string {
	var buf bytes.Buffer
	if r.Limit > 0 {
		buf.WriteString(fmt.Sprintf("&per_page=%d", r.Limit))
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
	if r.Lang != "" {
		buf.WriteString(fmt.Sprintf("&lang=%s", r.Lang))
	}
	if r.Lang == "" {
		buf.WriteString("&lang=zh")
	}
	if r.Orientation != "" {
		buf.WriteString(fmt.Sprintf("&orientation=%s", r.Orientation))
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
	if r.VideoType != "" && r.Type == configs.VideoType {
		buf.WriteString(fmt.Sprintf("&video_type=%s", r.VideoType))
	}
	return buf.String()
}

//获取API URL
func (r *RequestInfo) buildApiUrl() (apiUrl string) {
	params := r.buildParams()
	apiUrl = ImageApiUrl + "?key=" + r.ApiKey
	if r.Type == configs.VideoType {
		apiUrl = VideoApiUrl + "?key=" + r.ApiKey
	}
	r.ApiUrl = apiUrl + params
	return r.ApiUrl
}
