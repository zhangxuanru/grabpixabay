/*
@Time : 2020/6/3 16:33
@Author : zxr
@File : types
@Software: GoLand
*/
package spider

import (
	"context"
	"grabpixabay/app/scheduler"

	"github.com/PuerkitoBio/goquery"
)

//首页HTML结构体
type PixSearch struct {
	Html      *string
	Url       string
	Color     string
	Dom       *goquery.Document
	Ctx       context.Context
	Can       context.CancelFunc
	Scheduler *scheduler.SchedulPool
}
