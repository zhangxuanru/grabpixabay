package es

import (
	"context"
	"fmt"
	"grabpixabay/configs"
	"grabpixabay/core/api"
	"time"
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

func SavePicInfo(item api.ItemImage) {
	isHandpick := 0
	if item.EditorsChoice == true {
		isHandpick = 1
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
	response, e := client.Index().
		Index(configs.ES_INDEX).
		BodyJson(data).
		Do(context.Background())
	fmt.Println("err:", e)
	fmt.Println("resp:", response)
}
