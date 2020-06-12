/*
@Time : 2020/6/12 16:17
@Author : zxr
@File : storageimages
@Software: GoLand
*/
package scheduler

import (
	"fmt"
)

type Storage struct {
}

func newStorage() *Storage {
	return &Storage{}
}

//保存图片信息
func (s *Storage) SaveImages(item ItemImage) {

	fmt.Printf("SaveImages:%+v\n\n", item)
}
