/*
@Time : 2020/6/12 10:35
@Author : zxr
@File : db_test
@Software: GoLand
*/
package test

import (
	"grabpixabay/core/storage/models"
	"testing"
	"time"
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
