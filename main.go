/*
@Time : 2020/5/28 16:46
@Author : zxr
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"grabpixabay/app"
	"grabpixabay/configs"
)

//main   -type all     -page=10  -size=50  -color=all   全站抓取所有颜色图片，只抓10页,每页50条数据,
//main   -type latest  -page=10  -size=50  -color=red   抓取最新红色的图片，只抓10页, 每页50条数据

func main() {
	fmt.Printf("%v\n\n", configs.FULL_NAME)
	app.Run()
}
