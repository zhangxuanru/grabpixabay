package services

import (
	"path/filepath"
	"strings"
)

//获取图片类型
func GetImageType(typeStr string) int {
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
func GetImageFormat(path string) int {
	extMap := make(map[string]int)
	extMap["jpg"] = 1
	extMap["png"] = 2
	ext := strings.TrimLeft(filepath.Ext(path), ".")
	if format, ok := extMap[ext]; ok {
		return format
	}
	return 0
}
