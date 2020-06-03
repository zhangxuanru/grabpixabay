/*
@Time : 2020/5/29 17:00
@Author : zxr
@File : task
@Software: GoLand
*/
package distribute

import (
	"context"
	"fmt"
	"grabpixabay/app/crawler/pixabay"
	"grabpixabay/config"
)

type Task struct {
	TaskName  string
	Host      string
	Type      string
	PicUrl    string
	HostUrl   string
	Page      int
	StartTime int64
	Cxt       context.Context
	Can       context.CancelFunc
	CrawEngine
}

type CrawEngine struct {
	PxCrawler *pixabay.PixRequest
}

func NewTask() *Task {
	return &Task{}
}

//运行抓取任务
func (t *Task) RunTask() {
	fmt.Println(t.HostUrl, "开始抓取......")
	t.crawEngine()
	fmt.Println(t.HostUrl, "抓取结束......")
}

//判断任务是走哪个引擎，方便以后扩展抓取其它的网站
func (t *Task) crawEngine() {
	if t.Host == config.PIX_HOST {
		t.PxCrawler = &pixabay.PixRequest{
			HostUrl: t.HostUrl,
			PicUrl:  t.PicUrl,
			Page:    t.Page,
			Cxt:     t.Cxt,
			Can:     t.Can,
		}
		t.PxCrawler.CrawPixAbAyEngineType(t.Type)
		return
	}
	fmt.Printf("当前仅支持:%s的抓取", config.PIX_HOST)
	return
}
