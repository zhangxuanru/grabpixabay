/*
@Time : 2020/6/4 10:58
@Author : zxr
@File : types
@Software: GoLand
*/
package pixabay

import (
	"context"
	"grabpixabay/app/scheduler"
)

//发起抓取请求的结构体
type PixRequest struct {
	HostUrl string
	PicUrl  string
	Html    string
	Page    int
	Cxt     context.Context
	Can     context.CancelFunc
	SchPool scheduler.SchedulingPool
}

//全部抓取结构体
type CrawlerAll struct {
	Title      string
	PixRequest *PixRequest
	CurrPage   int
	VisitUrl   map[string]struct{} //记录访问过的URL，避免重复访问
	Worker     *worker
}
