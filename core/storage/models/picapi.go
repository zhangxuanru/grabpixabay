/*
@Time : 2020/6/12 14:56
@Author : zxr
@File : userstat
@Software: GoLand
*/
package models

import (
	"errors"
	"time"
)

//API日志表
type PicApi struct {
	Id      int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	PxImgId uint      `gorm:"unique; not null; comment:'px站的图片ID'" json:"px_img_id"`
	Api     string    `gorm:"type:text; not null; comment:'API数据'" json:"api"`
	AddTime time.Time `gorm:"comment:'添加时间'" json:"add_time"`
}

func NewPicApi() *PicApi {
	return &PicApi{}
}

func (p *PicApi) Save() (id int, err error) {
	if p.PxImgId == 0 {
		return 0, errors.New("pximgid is nil")
	}
	pic := &PicApi{}
	GetDB().Where("px_img_id = ?", p.PxImgId).Select("id").First(pic)
	if pic.Id > 0 {
		return pic.Id, nil
	}
	create := GetDB().Create(p)
	return p.Id, create.Error
}
