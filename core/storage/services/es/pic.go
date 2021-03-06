package es

import (
	"context"
	"grabpixabay/configs"
	"grabpixabay/core/api"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type ImageIndexData struct {
	PicId      int    `json:"pic_id"`
	UserId     int    `json:"user_id"`
	CategoryId int    `json:"category_id"`
	ImageType  string `json:"image_type"`
	PicColor   string `json:"pic_color"`
	Direction  string `json:"direction"`
	Category   string `json:"category"`
	Tags       string `json:"tags"`
	IsHandpick int    `json:"is_handpick"`
	AddDate    int64  `json:"add_date"`
}

//保存图片信息到ES中
func SavePicInfo(item api.ItemImage) {
	isHandpick := 0
	if item.EditorsChoice == true {
		isHandpick = 1
	}
	if item.ID == 0 || item.UserID == 0 {
		return
	}
	data := ImageIndexData{
		PicId:      item.ID,
		UserId:     item.UserID,
		CategoryId: 0,
		ImageType:  item.Type,
		PicColor:   item.Color,
		Direction:  item.Orientation,
		Category:   item.Category,
		Tags:       item.Tags,
		IsHandpick: isHandpick,
		AddDate:    time.Now().Unix(),
	}
	esId := strconv.Itoa(item.ID)
	_, e := client.Index().
		Index(configs.ES_INDEX).
		BodyJson(data).
		Id(esId).
		Do(context.Background())
	if e != nil {
		logrus.Error("SavePicInfo error :", e)
	}
}

//func EditEsPhoto(picId int) {
//	q := elastic.NewTermQuery("pic_id", picId)
//	//q := elastic.NewQueryStringQuery("pic_id:5183312")
//	res, err := client.DeleteByQuery().Index(configs.ES_INDEX).Query(q).Do(context.Background())
//	fmt.Printf("res:%+v", res)
//	fmt.Println("err:", err)
//}
