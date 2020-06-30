/*
@Time : 2020/6/11 11:54
@Author : zxr
@File : go_test
@Software: GoLand
*/
package bak

import (
	"fmt"

	"net/http"
	"testing"
)

func Test_Go(t *testing.T) {
	c := make(chan bool)
	b := make(chan string)
	go func() {
		for {
			select {
			case <-c:
				fmt.Println("true")
			case <-b:
				fmt.Println("false")
			}
		}
	}()
	fmt.Println("...........")
}

func TestUrlFile(t *testing.T) {
	url := "https://cdn.pixabay.com/photo/2020/06/06/14/36/ice-5266805__340.jpg"
	qiniu := &services.QiNiu{SrcFile: url}
	services.NewQiNiu()

	ret, e := qiniu.UploadFile()

	fmt.Printf("ret:%+v\n\n", ret)
	fmt.Printf("ret 2:%+v\n\n", *ret.PutRet)
	fmt.Println("err:", e)
	return

	name := qiniu.GenDefaultFileName()
	fmt.Println(name)
	return
	exists, err := qiniu.HeadExists()
	fmt.Println("exists", exists)
	fmt.Println("err:", err)

	request, _ := http.NewRequest(http.MethodHead, url, nil)
	response, err := http.DefaultClient.Do(request)
	//defer response.Body.Close()
	fmt.Println("response:", response)
	fmt.Println("err:", err)
}
