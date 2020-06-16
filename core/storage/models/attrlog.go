/*
@Time : 2020/6/12 17:32
@Author : zxr
@File : picture
@Software: GoLand
*/
package models

import (
	"errors"
	"time"
)

//图片属性日志表
type PictureAttrLog struct {
	Id       int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	PicId    uint      `gorm:"index:pic_id; not null; comment:'图片ID'" json:"pic_id"`
	ImageURL string    `gorm:"type:varchar(100); NOT NULL; comment:'源图片地址'" json:"image_url"`
	ErrMsg   string    `gorm:"type:varchar(100);   comment:'错误信息'" json:"err_msg"`
	AddTime  time.Time `gorm:"comment:'添加时间'" json:"add_time"`
}

func NewPictureAttrLog() *PictureAttrLog {
	return &PictureAttrLog{}
}

func (p *PictureAttrLog) Insert() (id int, err error) {
	if p.PicId == 0 || p.ErrMsg == "" {
		return 0, errors.New("picid or errmsg is nil")
	}
	create := GetDB().Create(p)
	return p.Id, create.Error
}
