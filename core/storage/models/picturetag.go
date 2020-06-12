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

//图片标签表
type PictureTag struct {
	Id         int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	PicId      uint      `gorm:"index:pic_id; not null; comment:'图片ID'" json:"pic_id"`
	TagId      uint      `gorm:" not null; comment:'标签ID'" json:"tag_id"`
	State      int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'状态 1:状态正常 0:删除'" json:"state"`
	AddTime    time.Time `gorm:"comment:'添加时间'" json:"add_time"`
	UpdateTime time.Time `gorm:"comment:'修改时间'" json:"update_time"`
}

func NewPictureTag() *PictureTag {
	GetDB().AutoMigrate(&PictureTag{})
	return &PictureTag{}
}

//插入数据
func (p *PictureTag) Insert() (id int, err error) {
	if p.PicId == 0 {
		return 0, errors.New("PicId is nil")
	}
	tmpTag := &PictureTag{}
	GetDB().Where("pic_id = ? AND tag_id = ?", p.PicId, p.TagId).Select("id").First(tmpTag)
	if tmpTag.Id > 0 {
		return tmpTag.Id, nil
	}
	create := GetDB().Create(p)
	return p.Id, create.Error
}
