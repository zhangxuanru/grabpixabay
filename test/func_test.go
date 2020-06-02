/*
@Time : 2020/6/1 18:10
@Author : zxr
@File : func_test
@Software: GoLand
*/
package test

import (
	"fmt"
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
