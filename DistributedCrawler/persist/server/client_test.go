package main

import (
	"GoLangIntro/DistributedCrawler/config"
	"GoLangIntro/DistributedCrawler/rpcsupport"
	"GoLangIntro/SingleThreadCrawler/engine"
	"GoLangIntro/SingleThreadCrawler/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Url: "http://newcar.xcar.com.cn/m49989/",
		Id:  "m49989",
		Payload: model.Car{
			Name:         "WEY VV72020款2.0T 豪华型",
			ImageURL:     "http://img1.xcarimg.com/PicLib/s/s12277_300.jpg",
			Size:         "4760×1931×1655mm",
			Fuel:         7.7,
			Transmission: "7挡双离合",
			Displacement: 2,
			MaxSpeed:     205,
			Acceleration: 8.5,
			Price:        16.98,
		},
	}

	result := ""
	client.Call(config.ItemSaverRpc, item, &result)

	if err != nil || result != "okay" {
		t.Errorf("Result: %s; err %s", result, err)
	}
}
