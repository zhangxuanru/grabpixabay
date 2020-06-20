/*
@Time : 2020/6/12 14:56
@Author : zxr
@File : userstat
@Software: GoLand
*/
package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//用户数据统计表
type UserStat struct {
	Id           int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	Uid          int64     `gorm:"unique; not null; comment:'用户ID'" json:"uid"`
	PxUid        int64     `gorm:"unique; not null; comment:'px 用户ID'" json:"px_uid"`
	PicNum       uint      `gorm:"type:INT(11); NOT NULL;  default:0;  comment:'图片总数'"  json:"pic_num"`
	ViewNum      uint      `gorm:"NOT NULL; default:0; comment:'总浏览次数'" json:"view_num"`
	DownloadsNum uint      `gorm:"type:INT(11); NOT NULL;  default:0;  comment:'下载量'"    json:"downloads_num"`
	LikeNum      uint      `gorm:"type:INT(11); NOT NULL;  default:0;  comment:'喜欢总量'"  json:"like_num"`
	CommentNum   uint      `gorm:"type:INT(11); NOT NULL;  default:0;  comment:'评论总量'"  json:"comment_num"`
	FollowerNum  uint      `gorm:"type:INT(11); NOT NULL;  default:0;  comment:'收藏总量'"  json:"follower_num"`
	AddTime      time.Time `gorm:"comment:'添加时间'" json:"add_time"`
	UpdateTime   time.Time `gorm:"comment:'修改时间'" json:"update_time"`
}

func NewUserStat() *UserStat {
	return &UserStat{}
}

//修改用户统计表
func (u *UserStat) UpdateStat() (affected int64, err error) {
	buildMap := map[string]interface{}{
		"pic_num":       gorm.Expr("pic_num + ?", u.PicNum),
		"view_num":      gorm.Expr("view_num + ?", u.ViewNum),
		"downloads_num": gorm.Expr("downloads_num + ?", u.DownloadsNum),
		"like_num":      gorm.Expr("like_num + ?", u.LikeNum),
		"comment_num":   gorm.Expr("comment_num + ?", u.CommentNum),
		"follower_num":  gorm.Expr("follower_num + ?", u.FollowerNum),
	}
	updates := GetDB().Model(u).Where("px_uid = ?", u.PxUid).Updates(buildMap).Omit("add_time")
	return updates.RowsAffected, updates.Error
}

//插入用户统计表
func (u *UserStat) Insert() (id int, err error) {
	create := GetDB().Create(u)
	return u.Id, create.Error
}

//根据UID查询ID，判断UID是否存在
func (u *UserStat) GetIdByUid(uid int) int {
	stat := &UserStat{}
	GetDB().Where("uid = ?", u.Uid).Select("id").First(stat)
	if stat.Id > 0 {
		return stat.Id
	}
	return 0
}
