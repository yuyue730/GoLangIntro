package persist

import (
	"GoLangIntro/SingleThreadCrawler/engine"
	"GoLangIntro/SingleThreadCrawler/model"
	"context"
	"encoding/json"
	"testing"

	"github.com/olivere/elastic/v7"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
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

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "car_test"
	err = save(client, expected, index)
	if err != nil {
		panic(err)
	}

	client, err = elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("car_profile").
		Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual engine.Item
	err = json.Unmarshal([]byte(resp.Source), &actual)
	actualCar, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualCar
	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("Got %v, expect %v", actual, expected)
	}
}
