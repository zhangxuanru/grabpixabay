/*
@Time : 2020/6/9 16:23
@Author : zxr
@File : test_json
@Software: GoLand
*/
package test

import (
	"fmt"
	"grabpixabay/core/scheduler"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func Test_Json(t *testing.T) {
	str := []byte(`{
		"total":1260102,
		"totalHits":500,
		"hits":[
	{
		"id":5255326,
		"pageURL":"https://pixabay.com/zh/photos/landscape-fantasy-sky-clouds-5255326/",
		"type":"photo",
		"tags":"\u666f\u89c2, \u5e7b\u60f3, \u5929\u7a7a",
		"previewURL":"https://cdn.pixabay.com/photo/2020/06/03/15/20/landscape-5255326_150.jpg",
		"previewWidth":150,
		"previewHeight":100,
		"webformatURL":"https://pixabay.com/get/53e2d0464950aa14f1dc8460962931771637deed5b4c704c7c2e7ddc9148c15f_640.jpg",
		"webformatWidth":640,
		"webformatHeight":427,
		"largeImageURL":"https://pixabay.com/get/53e2d0464950aa14f6da8c7dda7936781c3cd6e55b596c4870267ad29f4bc05cbf_1280.jpg",
		"imageWidth":7087,
		"imageHeight":4724,
		"imageSize":3912235,
		"views":21882,
		"downloads":18129,
		"favorites":64,
		"likes":137,
		"comments":84,
		"user_id":3764790,
		"user":"enriquelopezgarre",
		"userImageURL":"https://cdn.pixabay.com/user/2020/06/03/11-05-03-625_250x250.jpg"
	},
      {
			"id":5262901,
			"pageURL":"https://pixabay.com/zh/photos/doll-clown-sad-colorful-sweet-5262901/",
			"type":"photo",
			"tags":"\u5a03\u5a03, \u5c0f\u4e11, \u60b2\u4f24",
			"previewURL":"https://cdn.pixabay.com/photo/2020/06/05/11/27/doll-5262901_150.jpg",
			"previewWidth":150,
			"previewHeight":150,
			"webformatURL":"https://pixabay.com/get/53e2d3414352ad14f1dc8460962931771637deed5b4c704c7c2e7ddc9148c15f_640.jpg",
			"webformatWidth":640,
			"webformatHeight":640,
			"largeImageURL":"https://pixabay.com/get/53e2d3414352ad14f6da8c7dda7936781c3cd6e55b596c4870267ad29f4bc05cbf_1280.jpg",
			"imageWidth":6000,
			"imageHeight":6000,
			"imageSize":5800806,
			"views":4978,
			"downloads":4346,
			"favorites":29,
			"likes":42,
			"comments":9,
			"user_id":686414,
			"user":"Alexas_Fotos",
			"userImageURL":"https://cdn.pixabay.com/user/2020/05/01/11-54-53-871_250x250.png"
		}]
}`)

	image, _ := scheduler.ToApiImageResp(str)

	fmt.Printf("%+v", image)

}

func Test_Str(t *testing.T) {
	a := make([]rune, 100)
	for i := range a {
		a[i] = rune(RandInt(19968, 40869))
	}
	fmt.Println(string(a))
}

func RandInt(min, max int32) int32 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int31n(max-min)
}

func TestPic(t *testing.T) {
	privUrl := "https://cdn.pixabay.com/photo/2020/06/14/08/00/landscape-5296910_150.jpg"
	onePicSrc := strings.Replace(privUrl, "_150.", "__340.", 1)
	twoPicSrc := strings.Replace(privUrl, "_150.", "__480.", 1)
	threePicSrc := strings.Replace(privUrl, "_150.", "_960_720.", 1)

	fmt.Println(onePicSrc)
	fmt.Println(twoPicSrc)
	fmt.Println(threePicSrc)
}
