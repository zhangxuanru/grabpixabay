/*
@Time : 2020/5/27 18:12
@Author : zxr
@File : image2_test
@Software: GoLand
*/
package test

import (
	"fmt"
	"image"
	"log"
	"os"
	"testing"

	"github.com/EdlinOrg/prominentcolor"
)

//获取图片颜色
func TestImage2(t *testing.T) {

	// Step 1: Load the image
	imgUrl := "../static/images/fairytale-5213337_640.jpg"
	img, err := loadImage(imgUrl)
	if err != nil {
		log.Fatal("Failed to load image", err)
	}

	// Step 2: Process it
	colours, err := prominentcolor.Kmeans(img)
	if err != nil {
		log.Fatal("Failed to process image", err)
	}

	fmt.Println("Dominant colours:")
	for _, colour := range colours {
		fmt.Println("#" + colour.AsString())
	}
}

func loadImage(fileInput string) (image.Image, error) {
	f, err := os.Open(fileInput)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	return img, err
}
