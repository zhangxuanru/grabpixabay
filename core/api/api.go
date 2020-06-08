/*
@Time : 2020/6/8 18:42
@Author : zxr
@File : api
@Software: GoLand
*/
package api

import (
	"errors"
	"grabpixabay/configs"
)

//根据不同的类型返回不同的API
func NewApi(aType string) (Api, error) {
	switch aType {
	case configs.ImageType:
		return NewImageRequest(), nil
	case configs.VideoType:
		return NewVideoRequest(), nil
	}
	return nil, errors.New("type is Undefined")
}
