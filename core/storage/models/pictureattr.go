/*
@Time : 2020/6/12 17:32
@Author : zxr
@File : picture
@Software: GoLand
*/
package models

import "time"

//图片属性表，记录图片地址信息
type PictureAttr struct {
	Id         int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	PicId      uint      `gorm:"index:pic_id; not null; comment:'图片ID'" json:"pic_id"`
	ImageURL   string    `gorm:"type:varchar(100); NOT NULL; comment:'源图片地址'" json:"image_url"`
	Width      int       `gorm:"not null;default:0; comment:'图片宽度'" json:"width"`
	Height     int       `gorm:"not null;default:0; comment:'图片高度'" json:"height"`
	PicAddress string    `gorm:"type:varchar(100);   comment:'本地图片地址'" json:"pic_address"`
	FileName   string    `gorm:"type:varchar(255); NOT NULL; comment:'图片名称'" json:"file_name"`
	IsQiniu    int       `gorm:"type:TINYINT(1); NOT NULL;default:0; comment:'是否上传七牛 1:已上传 0:未上传'" json:"is_qiniu"`
	State      int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'状态 1:状态正常 0:删除'" json:"state"`
	AddTime    time.Time `gorm:"comment:'添加时间'" json:"add_time"`
	UpdateTime time.Time `gorm:"comment:'修改时间'" json:"update_time"`
}

func NewPictureAttr() *PictureAttr {
	return &PictureAttr{}
}

func (p *PictureAttr) Insert() (id int, err error) {
	create := GetDB().Create(p)
	return p.Id, create.Error
}

func (p *PictureAttr) EditUpload(upload int) (affected int64, err error) {
	update := GetDB().Model(&PictureAttr{}).Where("id = ?", p.Id).Update("is_qiniu", upload)
	return update.RowsAffected, update.Error
}

func (p *PictureAttr) GetIdByPicId() (pic *PictureAttr) {
	pic = &PictureAttr{}
	GetDB().Where("pic_id = ? AND width = ?", p.PicId, p.Width).Select("id,is_qiniu").First(pic)
	return pic
}
