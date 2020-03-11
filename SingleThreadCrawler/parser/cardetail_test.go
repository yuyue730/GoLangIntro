package parser

import (
	"GoLangIntro/SingleThreadCrawler/engine"
	"GoLangIntro/SingleThreadCrawler/model"
	"io/ioutil"
	"testing"
)

func TestParseCarDetail(t *testing.T) {
	contents, err := ioutil.ReadFile("./cardetail_test_data.html")

	if err != nil {
		panic(err)
	}

	results := ParseCarDetail(contents, "http://newcar.xcar.com.cn/m49989/")

	if len(results.Items) != 1 {
		t.Errorf("result.Items expect one element, however receive %d element(s)",
			len(results.Items))
	}

	expectItem := engine.Item{
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
		Type: "xcar",
		Id:   "m49989",
		Url:  "http://newcar.xcar.com.cn/m49989/",
	}
	actualItem := results.Items[0]

	if actualItem != expectItem {
		t.Errorf("actualItem != expectItem. actual=%+v expect=%+v",
			actualItem, expectItem)
	}

	expectRequestSize := 8
	if len(results.Requests) != expectRequestSize {
		t.Errorf("Expect requests size = %d, acutal requests size = %d",
			expectRequestSize, len(results.Requests))
	}

	expectUrls := []string{
		"http://newcar.xcar.com.cn/m45592/",
		"http://newcar.xcar.com.cn/m45592/",
		"http://newcar.xcar.com.cn/m46731/",
	}

	for i, url := range expectUrls {
		if results.Requests[i].Url != url {
			t.Errorf("index %d expect url %s, actual url %s",
				i, results.Requests[i].Url, url)
		}
	}
}