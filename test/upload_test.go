package test

import (
	"fmt"
	"grabpixabay/core/storage/models"
	"grabpixabay/core/storage/services"
	"testing"
)

func TestUploadUserHead(t *testing.T) {
	users := models.NewUser().GetList()
	for _, user := range users {
		if user.IsQiniu == 1 {
			continue
		}
		niu := services.QiNiu{
			SrcFile: user.HeadPortrait,
		}
		ret, err := niu.UploadFile()
		fmt.Printf("ret : %+v\n\n", ret.PutRet)
		if err != nil {
			fmt.Println("err:", err)
			continue
		}
		if ret.PutRet.Key != "" {
			u := &models.User{
				Id:       user.Id,
				FileName: ret.FileName,
				IsQiniu:  1,
			}
			u.UpdateUpload()
		}
	}
}

func TestAa(t *testing.T) {
	users := models.NewUser().GetList()
	for _, user := range users {
		niu := services.QiNiu{
			SrcFile: user.HeadPortrait,
		}
		fName := niu.GenDefaultFileName()
		u := &models.User{
			Id:       user.Id,
			FileName: fName,
			IsQiniu:  1,
		}
		u.UpdateUpload()
	}
}
