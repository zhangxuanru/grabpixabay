/*
@Time : 2020/6/5 14:51
@Author : zxr
@File : storage
@Software: GoLand
*/
package storage

import (
	"fmt"
)

var UA = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36"
var Referer = "https://pixabay.com/"

type Image struct {
	imageInfo  *ImageInfo
	uploadChan chan *ImageInfo
}

func NewStorageImage(image *ImageInfo) *Image {
	return &Image{
		imageInfo:  image,
		uploadChan: make(chan *ImageInfo, 1000),
	}
}

func (s *Image) Storage() {
	//写数据库

	//传完了图片，
	s.uploadImage()

	fmt.Println("Storage:", s.imageInfo)
}

func (s *Image) uploadImage() {
	//var (
	//	putResult *storage.PutRet
	//	err       error
	//)
	//upload := qiniu.QiNiu{
	//	SrcFile: s.imageInfo.ImgSrc,
	//	UA:      UA,
	//	Referer: Referer,
	//}
	//upload.UpFileName = upload.GenDefaultFileName()
	//if putResult, err = upload.UploadHttpFile(); err != nil || putResult.Key == "" {
	//	logrus.WithFields(logrus.Fields{
	//		"srcFile": upload.SrcFile,
	//	}).Error(err)
	//	s.uploadChan <- s.imageInfo
	//	return
	//}
}
