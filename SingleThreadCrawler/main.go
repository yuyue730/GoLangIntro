package main

import (
	"GoLangIntro/SingleThreadCrawler/engine"
	"GoLangIntro/SingleThreadCrawler/parser"
	"GoLangIntro/SingleThreadCrawler/persist"
	"GoLangIntro/SingleThreadCrawler/scheduler"
)

func main() {
	// Please use engine.SimpleEngine{} for single threaded web crawler
	// e := engine.SimpleEngine{}
	itemChan, err := persist.ItemSaver("car_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(
		engine.Request{
			Url: "http://newcar.xcar.com.cn/",
			Parser: engine.NewFuncParser(
				parser.ParseCarList, "ParseCarList"),
		})
}
