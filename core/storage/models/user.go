/*
@Time : 2020/6/11 17:57
@Author : zxr
@File : user
@Software: GoLand
*/
package models

import (
	"time"
)

//用户表
type User struct {
	Id           int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`                               //指定主键并自增
	Uid          int64     `gorm:"unique; not null; comment:'用户ID'" json:"uid"`                                         //用户ID
	NickName     string    `gorm:"type:varchar(100); NOT NULL; comment:'用户昵称'" json:"nick_name"`                        //昵称
	PersonAddr   string    `gorm:"type:varchar(100); NOT NULL; comment:'个人中心地址'" json:"person_addr"`                    //个人中心地址
	UserName     string    `gorm:"index:user_pass; type:varchar(50);  comment:'登录用户名'" json:"user_name"`                //登录用户名
	Passwd       string    `gorm:"index:user_pass; type:char(32);     comment:'登录密码'" json:"passwd"`                    //用户密码
	UserType     int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'用户类型 1:本站注册 2:px站抓取'" json:"user_type"` //用户类型 1,本站注册， 2：px站抓取
	HeadPortrait string    `gorm:"type:varchar(100); NOT NULL;comment:'头像地址'" json:"head_portrait"`                     //头像地址
	AddTime      time.Time `gorm:"comment:'添加时间'" json:"add_time"`                                                      //创建时间
	UpdateTime   time.Time `gorm:"comment:'修改时间'" json:"update_time"`                                                   //修改后自动更新时间
}

func NewUser() *User {
	GetDB().AutoMigrate(&User{})
	return &User{}
}

//插入数据，检查UID是否已存在， 不存在则插入，存在则直接返回
func (u *User) InsertCheckByUid() (id int, err error) {
	tmpUser := &User{}
	GetDB().Where("uid = ?", u.Uid).Select("id").First(tmpUser)
	if tmpUser.Id > 0 {
		return tmpUser.Id, nil
	}
	create := GetDB().Create(u)
	return u.Id, create.Error
}
