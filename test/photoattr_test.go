/*
@Time : 2020/6/30 11:45
@Author : zxr
@File : photoattr_test
@Software: GoLand
*/
package test

import (
	"fmt"
	"grabpixabay/core/api"
	"grabpixabay/core/storage/models"
	services "grabpixabay/core/storage/services/images"

	"testing"

	jsoniter "github.com/json-iterator/go"
)

//补充1280的图片
func TestAttrSize(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	list := models.NewPicture().GetList()
	for _, pic := range list {
		item := &api.ItemImage{}
		apiInfo := models.NewPicApi().GetApiById(int(pic.PxImgId))
		if len(apiInfo.Api) < 10 {
			continue
		}
		_ = json.Unmarshal([]byte(apiInfo.Api), item)
		newItem := *item
		service := &services.ImageService{}
		service.DownloadPic(newItem)
	}

}

//补充默认评论
func TestInsertComments(t *testing.T) {
	var defaultComments = []string{
		"very beautiful ",
	}
	fmt.Println(defaultComments)
	list := models.NewPicture().GetList()
	for _, item := range list {
		//for i := 0; i < 10; i++ {
		//	rand.Seed(time.Now().UnixNano())
		//	comment := models.NewComments()
		//	comment.PicId = int(item.PxUid)
		//	comment.AddTime = time.Now()
		//	comment.State = 1
		//	comment.Uid = rand.Intn(135)
		//	_, _ = comment.Insert()
		//}
		fmt.Println(item.Uid)
	}
}
