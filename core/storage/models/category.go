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

//图片分类表
type Category struct {
	Id           int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	CategoryName string    `gorm:"type:varchar(50); unique; NOT NULL; comment:'分类名'" json:"category_name"`
	State        int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'状态 1:状态正常 0:删除'" json:"state"`
	AddTime      time.Time `gorm:"comment:'添加时间'" json:"add_time"`
	UpdateTime   time.Time `gorm:"comment:'修改时间'" json:"update_time"`
}

func NewCategory() *Category {
	GetDB().AutoMigrate(&Category{})
	return &Category{}
}

//插入数据，
func (c *Category) Insert() (id int, err error) {
	if c.CategoryName == "" {
		return 0, errors.New("CategoryName is nil")
	}
	tmpCategory := &Category{}
	GetDB().Where("category_name = ?", c.CategoryName).Select("id").First(tmpCategory)
	if tmpCategory.Id > 0 {
		return tmpCategory.Id, nil
	}
	create := GetDB().Create(c)
	return c.Id, create.Error
}
