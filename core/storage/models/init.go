/*
@Time : 2020/6/11 18:06
@Author : zxr
@File : init
@Software: GoLand
*/
package models

import (
	"fmt"
	"grabpixabay/configs"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var Db *gorm.DB

//初始化MYSQL
func init() {
	var err error
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", configs.DbUser, configs.DbPassWd, configs.DbHost, configs.DbPort, configs.DbDataBase)
	if Db, err = gorm.Open("mysql", source); err != nil {
		logrus.Fatal("init mysql:", err)
		return
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "px_" + defaultTableName
	}

	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)

	Db.LogMode(true)
	Db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	//创建表
	//initMigrate()
}

//获取DB对象
func GetDB() *gorm.DB {
	return Db
}

//创建表
func initMigrate() {
	GetDB().AutoMigrate(&User{}, &UserStat{}, &Tag{}, &PictureTag{}, &PictureAttr{}, &Picture{}, &Category{}, &PicApi{})
}
