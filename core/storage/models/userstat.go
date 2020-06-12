/*
@Time : 2020/6/12 14:56
@Author : zxr
@File : userstat
@Software: GoLand
*/
package models

import "time"

//用户数据统计表
type UserStat struct {
	Id           int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	Uid          int64     `gorm:"unique; not null; comment:'用户ID'" json:"uid"`
	PicNum       uint      `gorm:"type:INT(11); NOT NULL;  default:0;  comment:'图片总数'"  json:"pic_num"`
	DownloadsNum uint      `gorm:"type:INT(11); NOT NULL;  default:0;  comment:'下载量'"    json:"downloads_num"`
	LikeNum      uint      `gorm:"type:INT(11); NOT NULL;  default:0;  comment:'喜欢总量'"  json:"like_num"`
	CommentNum   uint      `gorm:"type:INT(11); NOT NULL;  default:0;  comment:'评论总量'"  json:"comment_num"`
	FollowerNum  uint      `gorm:"type:INT(11); NOT NULL;  default:0;  comment:'收藏总量'"  json:"follower_num"`
	AddTime      time.Time `gorm:"comment:'添加时间'" json:"add_time"`
	UpdateTime   time.Time `gorm:"comment:'修改时间'" json:"update_time"`
}

func NewUserStat() *UserStat {
	GetDB().AutoMigrate(&UserStat{})
	return &UserStat{}
}
