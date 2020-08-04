/*
@Time : 2020/7/24 14:37
@Author : zxr
@File : product_es_test
@Software: GoLand
*/
package test

import (
	"fmt"
	"grabpixabay/configs"
	"grabpixabay/core/api"
	"grabpixabay/core/storage/models"
	"grabpixabay/core/storage/services/es"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//将图片数据填充到ES中
func TestSaveEsImages(t *testing.T) {
	configs.AppConfig()

	list := models.NewPicture().GetList()
	for _, item := range list {
		data := models.NewPicApi().GetApiById(int(item.PxImgId))
		apiText := data.Api

		resp := &api.ItemImage{}
		if err := json.Unmarshal([]byte(apiText), resp); err != nil {
			fmt.Println("err:", err)
			continue
		}
		es.SavePicInfo(*resp)
	}
	fmt.Println("OK")
}
