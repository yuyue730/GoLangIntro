package parser

import (
	"GoLangIntro/SingleThreadCrawler/engine"
	"regexp"
)

var carDetailRegexp = `<a href="(/m\d+/)" target="_blank"`

// This function goes over a specific Car Model Page and fetch all specific car
// details link whose href="/m\d+/"
func ParseCarModel(contents []byte) engine.ParseResult {
	regexpPtrDetail := regexp.MustCompile(carDetailRegexp)
	matchesDetail := regexpPtrDetail.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matchesDetail {
		result.Requests = append(result.Requests, engine.Request{
			Url:        host + string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}

	return result
}
