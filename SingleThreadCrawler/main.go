package main

import (
	"GoLangIntro/SingleThreadCrawler/engine"
	"GoLangIntro/SingleThreadCrawler/parser"
)

func main() {
	engine.Run(
		engine.Request{
			Url: "http://newcar.xcar.com.cn/",
			Parser: engine.NewFuncParser(
				parser.ParseCarList, "ParseCarList"),
		})
}
