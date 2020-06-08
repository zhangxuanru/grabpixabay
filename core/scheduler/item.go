/*
@Time : 2020/6/8 15:52
@Author : zxr
@File : request
@Software: GoLand
*/
package scheduler

import (
	"fmt"
	"grabpixabay/core/api"

	"github.com/sirupsen/logrus"
)

func NewItem() *Item {
	return &Item{}
}

//发起请求
func (i *Item) Start() {

	if request, err := api.NewApi(i.Command.Type); err != nil {
		logrus.Error(err)
		return
	}

	fmt.Printf("%+v", *i.Command)
}

func (i *Item) SetParams() {
	request := api.NewImageRequest()
	request.Type = i.Command.Type
	request.Size = i.Command.Size
	request.Order = i.Command.Order
	request.Color = i.Command.Color
}

//监听信号
func (i *Item) Monitor() {
	select {
	case <-i.SignChan:
		i.Can()
		fmt.Println("end ####")
	case <-i.Ctx.Done():
		fmt.Println("end ctx ....")
		return
	}
}
