/*
@Time : 2020/6/12 17:32
@Author : zxr
@File : picture
@Software: GoLand
*/
package models

import "time"

//图片主表
type Picture struct {
	Id             int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	Uid            int64     `gorm:"not null; comment:'用户ID'" json:"uid"`
	PxUid          int64     `gorm:"not null; comment:'px 用户ID'" json:"px_uid"`
	PxImgId        uint      `gorm:"unique; not null; comment:'px站的图片ID'" json:"px_img_id"`
	PageURL        string    `gorm:"type:varchar(100); NOT NULL; comment:'网页页面地址'" json:"page_url"`
	CategoryId     uint      `gorm:"index:category_id; not null; comment:'分类ID'" json:"category_id"`
	ImageFormat    int       `gorm:"type:TINYINT(1); NOT NULL;default:0; comment:'图片格式 1:jpg 2:png 0:其它'" json:"image_format"`
	ThemeColor     string    `gorm:"type:varchar(20); comment:'主体颜色'" json:"theme_color"`
	ImageDirection int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'图像方向 1:未区分 2:水平 3:垂直'" json:"image_direction"`
	ImageType      int       `gorm:"type:TINYINT(1); NOT NULL;default:0; comment:'图片类型 0:未区分 1:照片 2:插图 3:矢量'" json:"image_type"`
	ViewNum        uint      `gorm:"NOT NULL; default:0; comment:'总浏览次数'" json:"view_num"`
	DownloadsNum   uint      `gorm:"NOT NULL; default:0; comment:'下载总数'" json:"downloads_num"`
	LikeNum        uint      `gorm:"NOT NULL; default:0; comment:'喜欢总数'" json:"like_num"`
	FavoritesNum   uint      `gorm:"NOT NULL; default:0; comment:'收藏总数'" json:"favorites_num"`
	CommentsNum    uint      `gorm:"NOT NULL; default:0; comment:'评论总数'" json:"comments_num"`
	IsHandpick     int       `gorm:"type:TINYINT(1); NOT NULL;default:0; comment:'是否精选 1:精选 0:不是精选'" json:"is_hand_pick"`
	State          int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'状态 1:状态正常 0:删除'" json:"state"`
	AddTime        time.Time `gorm:"comment:'添加时间'" json:"add_time"`
	UpdateTime     time.Time `gorm:"comment:'修改时间'" json:"update_time"`
}

func NewPicture() *Picture {
	GetDB().AutoMigrate(&Picture{})
	return &Picture{}
}

func (p *Picture) Save() (id int, err error) {
	pic := &Picture{}
	GetDB().Where("px_img_id = ?", p.PxImgId).Select("id").First(pic)
	if pic.Id > 0 {
		return pic.Id, nil
	}
	create := GetDB().Create(p)
	return p.Id, create.Error
}
