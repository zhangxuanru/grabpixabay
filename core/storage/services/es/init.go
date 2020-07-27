package es

import (
	"context"
	"fmt"
	"grabpixabay/configs"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var client *elastic.Client
var host = configs.ES_HOST

func init() {
	log := logrus.New()
	var err error
	client, err = elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetErrorLog(log),
		elastic.SetURL(host),
		elastic.SetBasicAuth(configs.ESUSER, configs.ESPASSWORD),
	)
	if err != nil {
		logrus.Error("elastic.NewClient err:", err)
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		logrus.Error("client.Ping err:", err)
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	esversion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
}
