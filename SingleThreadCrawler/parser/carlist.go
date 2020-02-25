package parser

import (
	"GoLangIntro/SingleThreadCrawler/engine"
	"regexp"
)

var host = "http://newcar.xcar.com.cn"
var carModelRegex = `<a href="(/\d+/)" target="_blank" class="list_img">`
var otherCarListRegex = `(//newcar.xcar.com.cn/car/[\d+-]+\d+)/`

// This function starts to crawl for all car informations. It will look at two
// parts of http://newcar.xcar.com.cn.
// Part I: Car Model already show on this page, which will match `carModelRegex`
// Part II: Additional Car List, which will match `otherCarListRegex`
func ParseCarList(contents []byte) engine.ParseResult {
	// Part I: Car Model already show on this page, which will match `carModelRegex`
	regexpPtrModel := regexp.MustCompile(carModelRegex)
	matchesModel := regexpPtrModel.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matchesModel {
		result.Requests = append(result.Requests, engine.Request{
			Url: host + string(m[1]),
			// This will be `http://newcar.xcar.com.cn/[id]`
			ParserFunc: ParseCarModel,
		})
	}

	// Part II: Additional Car List, which will match `otherCarListRegex`
	regexpPtrOtherList := regexp.MustCompile(otherCarListRegex)
	matchesOtherList := regexpPtrOtherList.FindAllSubmatch(contents, -1)

	for _, m := range matchesOtherList {
		result.Requests = append(result.Requests, engine.Request{
			Url: "http:" + string(m[1]),
			// This will be something like `http://newcar.xcar.com.cn/0-1-2-3-4-5`
			ParserFunc: ParseCarList,
		})
	}

	return result
}
