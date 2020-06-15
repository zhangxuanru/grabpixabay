package services

import (
	"grabpixabay/core/storage/models"
	"path/filepath"
	"strings"
)

//获取图片类型
func (i *ImageService) GetImageType(typeStr string) int {
	typeMap := make(map[string]int)
	typeMap["photo"] = 1
	typeMap["illustration"] = 2
	typeMap["vector"] = 3
	if imgType, ok := typeMap[typeStr]; ok {
		return imgType
	}
	return 0
}

//获取图片类型
func (i *ImageService) GetImageFormat(path string) int {
	extMap := make(map[string]int)
	extMap["jpg"] = 1
	extMap["png"] = 2
	ext := strings.TrimLeft(filepath.Ext(path), ".")
	if format, ok := extMap[ext]; ok {
		return format
	}
	return 0
}

//根据图片作者ID获取真实的用户ID
func (i *ImageService) GetUidByAuthorId(authorId int) int {
	if id, ok := i.AuthorMap[authorId]; ok {
		return id
	}
	if id := models.NewUser().GetUidByAuthorId(authorId); id > 0 {
		return id
	}
	return 0
}
