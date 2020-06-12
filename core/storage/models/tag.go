/*
@Time : 2020/6/11 17:57
@Author : zxr
@File : user
@Software: GoLand
*/
package models

import (
	"errors"
	"time"
)

//标签表
type Tag struct {
	Id         int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	TagName    string    `gorm:"type:varchar(50); unique; NOT NULL; comment:'标签名'" json:"tag_name"`
	State      int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'状态 1:状态正常 0:删除'" json:"state"`
	AddTime    time.Time `gorm:"comment:'添加时间'" json:"add_time"`
	UpdateTime time.Time `gorm:"comment:'修改时间'" json:"update_time"`
}

func NewTag() *Tag {
	GetDB().AutoMigrate(&Tag{})
	return &Tag{}
}

//插入数据
func (t *Tag) Insert() (id int, err error) {
	if t.TagName == "" {
		return 0, errors.New("TagName is nil")
	}
	tmpTag := &Tag{}
	GetDB().Where("tag_name = ?", t.TagName).Select("id").First(tmpTag)
	if tmpTag.Id > 0 {
		return tmpTag.Id, nil
	}
	create := GetDB().Create(t)
	return t.Id, create.Error
}
