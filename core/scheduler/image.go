/*
@Time : 2020/6/9 17:03
@Author : zxr
@File : image
@Software: GoLand
*/
package scheduler

import (
	"fmt"
	"grabpixabay/configs"
	"grabpixabay/core/api"
	"math"

	"github.com/sirupsen/logrus"
)

//拼凑各种查询条件，抓取所有图片
func (i *Item) CallImageAll() {
	i.callImageCmdDefault()
	if i.Command.More == true {
		i.callImageEditors()
		i.callImageOrientation()
		i.callImageCategory()
		if i.Command.Order == configs.OrderPopular {
			i.callImageOrderLast()
		}
		if i.Command.Color == "" {
			i.callImageAllColor()
		}
		if i.Command.Color == "" && i.Command.Order == configs.OrderPopular {
			i.callImageAllColorLast()
		}
		if i.Command.ImgType == configs.All {
			i.callImageType()
		}
		//随机生成查询字符串
		i.CallImageQuery("")
	}

}

//按默认命令行参数抓取图片
func (i *Item) callImageCmdDefault() {
	var (
		reqObj *api.RequestInfo
	)
	reqObj = i.getRequest()
	i.unifyRequest(reqObj)
}

//按最新的排序抓取图片
func (i *Item) callImageOrderLast() {
	var (
		reqObj *api.RequestInfo
	)
	reqObj = i.getRequest()
	reqObj.Order = configs.OrderLatest
	i.unifyRequest(reqObj)
}

//按所有颜色抓取图片
func (i *Item) callImageAllColor() {
	reqObj := i.getRequest()
	for _, color := range configs.GConf.Colors {
		reqObj.Color = color
		i.unifyRequest(reqObj)
	}
}

//按所有颜色并按最新的排序 来抓取图片
func (i *Item) callImageAllColorLast() {
	reqObj := i.getRequest()
	reqObj.Order = configs.OrderLatest
	for _, color := range configs.GConf.Colors {
		reqObj.Color = color
		i.unifyRequest(reqObj)
	}
}

//按图片类型抓取图片
func (i *Item) callImageType() {
	reqObj := i.getRequest()
	for _, imgType := range configs.GConf.ImageType {
		reqObj.ImageType = imgType
		i.unifyRequest(reqObj)
	}
}

//选择已获得编辑选择奖的图像
func (i *Item) callImageEditors() {
	reqObj := i.getRequest()
	reqObj.EditorsChoice = true
	i.unifyRequest(reqObj)
}

//图像宽于高还是宽于高 todo
func (i *Item) callImageOrientation() {
	fmt.Println("all...")
}

//按不同的分类抓取图片
func (i *Item) callImageCategory() {
	reqObj := i.getRequest()
	for _, category := range configs.GConf.Category {
		reqObj.Category = category
		i.unifyRequest(reqObj)
	}
}

//按搜索关键字抓取图片
func (i *Item) CallImageQuery(search string) {
	req := &api.RequestInfo{
		Q:      search,
		Type:   configs.ImageType,
		ApiKey: configs.ApiKey,
	}
	i.unifyRequest(req)
}

//如果设置了抓取总页数 按总页数抓取图片
func (i *Item) CallImageTotalPage() {
	var (
		apiResp   *ApiImageResp
		err       error
		totalPage int //总页
		reqObj    *api.RequestInfo
	)
	reqObj = i.getRequest()
	totalPage = i.Command.CountPage
	isTotalPage := false
	if totalPage > 0 {
		for j := 1; j <= totalPage; j++ {
			reqObj.Page = j
			if apiResp, err = i.distributeImage(reqObj); err != nil {
				break
			}
			if isTotalPage == false {
				//判断返回的总数是否能满足设置的总页数
				tmpTotal := int(math.Ceil(float64(apiResp.TotalHits) / float64(i.Command.Size)))
				if totalPage > tmpTotal {
					totalPage = tmpTotal
				}
				isTotalPage = true
			}
		}
	} else {
		fmt.Println("command CountPage eq 0")
	}
	return
}

//抓取所有图片时，统一的请求方法
func (i *Item) unifyRequest(reqObj *api.RequestInfo) {
	var (
		apiResp   *ApiImageResp
		totalPage int //总页
		err       error
	)
	if apiResp, err = i.distributeImage(reqObj); err != nil {
		return
	}
	totalPage = int(math.Ceil(float64(apiResp.TotalHits) / float64(i.Command.Size)))
	for j := 2; j <= totalPage; j++ {
		reqObj.Page = j
		if apiResp, err = i.distributeImage(reqObj); err != nil {
			break
		}
	}
}

//分用的分发图片的item
func (i *Item) distributeImage(reqObj *api.RequestInfo) (apiResp *ApiImageResp, err error) {
	var (
		resp []byte
	)
	if resp, err = reqObj.RequestImage(); err != nil {
		logrus.Error("RequestImage error:", err)
		return nil, err
	}
	if apiResp, err = ToApiImageResp(resp); err != nil {
		logrus.Error("ToApiImageResp error:", err)
		return nil, err
	}
	//分发item

	//
	return apiResp, nil
}
