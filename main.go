/*
@Time : 2020/5/28 16:46
@Author : zxr
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"grabpixabay/config"
	"grabpixabay/exec"
)

//main -host pixabay  -type all     -page=10     全站抓取，只抓10页
//main -host pixabay  -type latest  -page=10     获取最新  只抓10页
//main -host pixabay  -type sift      -page=10     获取小编精选  只抓10页
//main -host pixabay  -type pic  https://pixabay.com/zh/photos/fiber-cable-connection-network-4814456/ 获取图片

func init() {
	config.AppConfig()
}

func main() {
	fmt.Printf("%v\n\n", config.FULL_NAME)
	exec.Run()
}
