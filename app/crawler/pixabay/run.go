/*
@Time : 2020/6/1 11:58
@Author : zxr
@File : crawler
@Software: GoLand
*/
package pixabay

import (
	"fmt"
	"grabpixabay/config"
)

type PixRequest struct {
	HostUrl string
	PicUrl  string
	Html    string
	Page    int
}

func NewPixRequest() *PixRequest {
	return &PixRequest{}
}

//判断抓取的类型
func (p *PixRequest) CrawPixAbAyEngineType(crawType string) {
	switch crawType {
	case config.TYPE_ALL:
		p.RunAll()
	case config.TYPE_LATEST:
		p.RunLatest()
	case config.TYPE_SIFT:
		p.RunSift()
	case config.TYPE_PIC:
		p.RunPic()
	default:
		fmt.Println("type is Undefined")
		return
	}
}

//执行全站图片所有抓取
func (p *PixRequest) RunAll() {
	NewCrawlerAll(p).CrawlerAll()
}

func (p *PixRequest) RunLatest() {
	fmt.Println("latest 待开发....")
}

func (p *PixRequest) RunSift() {
	fmt.Println("sift 待开发....")
}

func (p *PixRequest) RunPic() {
	fmt.Println("pic 待开发....")
}
