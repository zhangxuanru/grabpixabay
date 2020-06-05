/*
@Time : 2020/6/1 18:10
@Author : zxr
@File : func_test
@Software: GoLand
*/
package test

import (
	"fmt"
	"grabpixabay/common/qiniu"
	"path"
	"testing"
)

type reqFunc func(url string, age int) int
type reqFunc2 func(url string, age int) string

func TA(req reqFunc2) reqFunc {
	return func(url string, age int) int {
		req(url, age)
		return 0
	}
}

//
//func Tb(f1 reqFunc) reqFunc {
//	return func(url string, age int) int {
//		f1(url, age)
//		return 1
//	}
//}

func Test_abc(t *testing.T) {

	c := TA(func(url string, age int) string {
		fmt.Println("-----")
		return ""
	})
	c("99,", 11)
	//c := TA(func(url string, age int) int {
	//	fmt.Println("url:", url, "age:", age)
	//	return 1
	//})
	//dd := c("aa", 12)
	//fmt.Println(dd)
}

func TestFile(t *testing.T) {
	var UA = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36"
	var Referer = "https://pixabay.com/"
	srcFile := "https://cdn.pixabay.com/photo/2020/05/23/17/52/strawberries-5210753_960_720.jpg"
	srcFile = "b.jpg"
	upload := qiniu.QiNiu{
		UA:         UA,
		Referer:    Referer,
		SrcFile:    srcFile,
		UpFileName: path.Base(srcFile),
	}
	err, ret := upload.UploadHttpFile()
	fmt.Println(err)
	fmt.Println("-------")
	fmt.Println(ret)
}
