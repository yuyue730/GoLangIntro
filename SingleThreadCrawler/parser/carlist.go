package parser

import (
	"GoLangIntro/SingleThreadCrawler/engine"
	"regexp"
)

var host = "http://newcar.xcar.com.cn"
var carModelRe = `<a href="(/\d+/)" target="_blank" class="list_img">`

func ParseCarList(contents []byte) engine.ParseResult {
	// Process already existed car models on main page
	reModel := regexp.MustCompile(carModelRe)
	matchesModel := reModel.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matchesModel {
		result.Requests = append(result.Requests, engine.Request{
			Url:        host + string(m[1]), // This will be `http://newcar.xcar.com.cn/[id]`
			ParserFunc: engine.NilParser,
		})
	}

	return result
}
