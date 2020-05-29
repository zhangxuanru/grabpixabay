/*
@Time : 2020/5/27 14:51
@Author : zxr
@File : image1_test
@Software: GoLand
*/
package test

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"testing"
)

func TestImg1(t *testing.T) {
	imagePath := "../static/images/fairytale-5213337_960_720.webp"
	file, _ := os.Open(imagePath)
	c, _, _ := image.DecodeConfig(file)

	fmt.Println("width = ", c.Width)
	fmt.Println("height = ", c.Height)

}
