/*
@Time : 2020/6/12 17:32
@Author : zxr
@File : picture
@Software: GoLand
*/
package models

import (
	"time"
)

//图片评论表
type Comments struct {
	Id      int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	PicId   int       `gorm:"index:pic_id; not null; comment:'图片ID'" json:"pic_id"`
	content string    `gorm:"type:text; NOT NULL; comment:'评论内容'" json:"content"`
	Uid     int       `gorm:"index:uid;   comment:'用户ID'" json:"uid"`
	State   int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'状态 1:状态正常 0:删除'" json:"state"`
	AddTime time.Time `gorm:"comment:'添加时间'" json:"add_time"`
}

func NewComments() *Comments {
	return &Comments{}
}

func (c *Comments) Insert() (id int, err error) {
	create := GetDB().Create(c)
	return c.Id, create.Error
}
