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

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/xormplus/xorm"
)

var Db *xorm.Engine

//初始化MYSQL
func init() {
	var err error
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", configs.DbUser, configs.DbPassWd, configs.DbHost, configs.DbPort, configs.DbDataBase)
	if Db, err = xorm.NewEngine("mysql", source); err != nil {
		logrus.Fatal("init mysql:", err)
		return
	}
	Db.ShowSQL(true)
	Db.SetLogger(logrus.DebugLevel)
	Db.SetMaxIdleConns(400)
	Db.SetMaxOpenConns(300)
}
