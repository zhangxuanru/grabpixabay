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
	"testing"
)

//更新分类ID数据到图片表中
func TestEditPicCategoryId(t *testing.T) {
	configs.AppConfig()
	var (
		categoryId int
		ok         bool
	)

	categoryMap := make(map[string]int)
	list := models.NewPicture().GetList()
	for _, item := range list {
		if item.Id < 1100 {
			continue
		}
		pxId := int(item.PxImgId)
		data := models.NewPicApi().GetApiById(pxId)
		resp := &api.ItemImage{}
		if err := json.Unmarshal([]byte(data.Api), resp); err != nil {
			fmt.Println("err:", err)
			continue
		}
		if resp.Category == "" {
			fmt.Println("category is nil")
			continue
		}

		if categoryId, ok = categoryMap[resp.Category]; !ok {
			categoryId = models.NewCategory().GetIdByCateName(resp.Category)
			categoryMap[resp.Category] = categoryId
		}

		if categoryId > 0 {
			if _, err := models.NewPicture().EditCategoryId(pxId, categoryId); err != nil {
				fmt.Println(pxId, "更新分类ID失败 err:", err)
			}
		}
	}
	fmt.Println("OK")
}
