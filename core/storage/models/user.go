/*
@Time : 2020/6/11 17:57
@Author : zxr
@File : user
@Software: GoLand
*/
package models

import "github.com/sirupsen/logrus"

//定义结构体
type User struct {
	Id            int    `xorm:"pk autoincr comment('自增ID')"`                        //指定主键并自增
	Uid           int64  `xorm:"unique not null comment('用户ID') "`                   //用户ID
	Nick_name     string `xorm:"varchar(100) not null comment('用户昵称') "`             //昵称
	Person_addr   string `xorm:"varchar(100) not null comment('个人中心地址')"`            //个人中心地址
	User_name     string `xorm:"index(user_pass) not null comment('登录用户名')"`         //登录用户名
	Passwd        string `xorm:"index(user_pass) char(32) not null comment('登录密码')"` //用户密码
	User_type     int    `xorm:"TINYINT(1) not null comment('用户类型 1:本站注册 2:px站抓取')"` //用户类型 1,本站注册， 2：px站抓取
	Head_portrait string `xorm:"varchar(100) comment('头像地址')"`                       //头像地址
	Add_time      int64  `xorm:"created comment('创建时间')"`                            //创建时间
	Update_time   int64  `xorm:"updated comment('更新时间')"`                            //修改后自动更新时间
}

func Sync2User() error {
	if err := Db.Sync2(new(User)); err != nil {
		logrus.Infoln("数据表同步失败 error:", err)
		return err
	}
	return nil
}

func NewUserModel() *User {
	return &User{}
}
