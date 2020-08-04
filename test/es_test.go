package test

import (
	"grabpixabay/configs"
	"grabpixabay/core/api"
	"grabpixabay/core/storage/services/es"
	"testing"
)

func TestSavePic(t *testing.T) {
	configs.AppConfig()
	item := api.ItemImage{
		ID:   5296910,
		Type: "photo",
		Tags: "景观, 沙漠, 沙丘",
		ItemReqParams: api.ItemReqParams{
			Category:      "nature",
			Orientation:   "horizontal",
			EditorsChoice: true,
		},
		ItemAuthor: api.ItemAuthor{
			UserID: 3764790,
		},
	}
	es.SavePicInfo(item)
}
