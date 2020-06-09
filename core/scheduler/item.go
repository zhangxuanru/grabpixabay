/*
@Time : 2020/6/8 15:52
@Author : zxr
@File : request
@Software: GoLand
*/
package scheduler

import (
	"fmt"
	"grabpixabay/configs"
	"grabpixabay/core/api"
)

//请求图片API
func (i *Item) CallImage() {
	if i.Command.CountPage > 0 {
		i.CallImageTotalPage()
		return
	}
	// 如果没设置总页数，则拼凑各种查询条件再抓取更多的图片
	//【因为一个查询最多只返回500条数据，尽可能多的抓取就先暂时拼凑各种查询条件】
	i.CallImageAll()
	return
}

//请求视频API
func (i *Item) CallVideo() {
	fmt.Println("抓取视频 暂时不支持, 下期开发....")
	return
}

//根据命令行参数构建request结构
func (i *Item) getRequest() *api.RequestInfo {
	request := &api.RequestInfo{
		Type:      i.Command.Type,
		Limit:     i.Command.Size,
		Order:     i.Command.Order,
		Color:     i.Command.Color,
		ImageType: i.Command.ImgType,
		VideoType: i.Command.VideoType,
		Q:         i.Command.Query,
		ApiKey:    configs.ApiKey,
		Page:      1,
		Lang:      "zh",
	}
	return request
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
