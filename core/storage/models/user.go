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

const UserPx = 2
const UserStation = 1

//用户表
type User struct {
	Id           int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`                               //指定主键并自增
	PxUid        int64     `gorm:"unique; not null; comment:'px 用户ID'" json:"px_uid"`                                   //用户ID
	NickName     string    `gorm:"type:varchar(100); NOT NULL; comment:'用户昵称'" json:"nick_name"`                        //昵称
	UserName     string    `gorm:"index:user_pass; type:varchar(50);  comment:'登录用户名'" json:"user_name"`                //登录用户名
	Passwd       string    `gorm:"index:user_pass; type:char(32);     comment:'登录密码'" json:"passwd"`                    //用户密码
	UserType     int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'用户类型 1:本站注册 2:px站抓取'" json:"user_type"` //用户类型 1,本站注册， 2：px站抓取
	HeadPortrait string    `gorm:"type:varchar(100); NOT NULL;comment:'头像地址'" json:"head_portrait"`                     //头像地址
	FileName     string    `gorm:"type:varchar(255); NOT NULL; comment:'图片名称'" json:"file_name"`
	IsQiniu      int       `gorm:"type:TINYINT(1); NOT NULL;default:0; comment:'是否上传七牛 1:已上传 0:未上传'" json:"is_qiniu"`
	AddTime      time.Time `gorm:"comment:'添加时间'" json:"add_time"`    //创建时间
	UpdateTime   time.Time `gorm:"comment:'修改时间'" json:"update_time"` //修改后自动更新时间
}

func NewUser() *User {
	return &User{}
}

//插入数据，检查UID是否已存在， 不存在则插入，存在则直接返回
func (u *User) InsertCheckByUid() (id int, err error) {
	tmpUser := &User{}
	GetDB().Where("px_uid = ?", u.PxUid).Select("id").First(tmpUser)
	if tmpUser.Id > 0 {
		return tmpUser.Id, nil
	}
	create := GetDB().Create(u)
	return u.Id, create.Error
}

//根据作者ID 查询用户ID
func (u *User) GetUidByAuthorId(authorId int) int {
	tmpUser := &User{}
	GetDB().Where("px_uid = ?", authorId).Select("id").First(tmpUser)
	if tmpUser.Id > 0 {
		return tmpUser.Id
	}
	return 0
}
