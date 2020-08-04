/*
@Time : 2020/6/12 10:35
@Author : zxr
@File : db_test
@Software: GoLand
*/
package test

import (
	"fmt"
	"grabpixabay/configs"
	"grabpixabay/core/storage/models"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

//用户主表
func TestUserDB(t *testing.T) {
	user := models.NewUser()
	user.PxUid = 29
	user.NickName = "zxr"
	user.UserType = 1
	user.HeadPortrait = "www.13520v.com"
	user.AddTime = time.Now()
	user.UpdateTime = time.Now()

	user.InsertCheckByUid()
}

//用户统计表
func TestUserStatDb(t *testing.T) {

	models.NewUserStat()
}

func TestImageTag(t *testing.T) {
	//保存图片tag信息
	picTag := models.NewPictureTag()
	picTag.PicId = uint(5255326)
	picTag.TagId = "1,2,3"
	picTag.State = models.StatusDefault
	picTag.AddTime = time.Now()
	picTag.UpdateTime = time.Now()
	now := time.Now()
	if id, err := picTag.Insert(); id > 0 {
		fmt.Println("id:", id)
	} else {
		logrus.Error("picTag.Insert error :", err)
	}
	since := time.Since(now)
	fmt.Println(since.String())
}

func TestTag(t *testing.T) {
	tag := models.NewTag()
	tag.TagName = "景观"
	tag.State = 1
	tag.AddTime = time.Now()
	tag.UpdateTime = time.Now()
	id, err := tag.Insert()
	fmt.Println("id=", id, "err=", err)
}

func TestCategory(t *testing.T) {
	configs.AppConfig()
	for _, category := range configs.GConf.Category {
		cate := models.Category{
			CategoryName: category,
			State:        1,
			AddTime:      time.Now(),
			UpdateTime:   time.Now(),
		}
		id, err := cate.Insert()
		fmt.Println("id:", id)
		fmt.Println("err:", err)
	}
}
