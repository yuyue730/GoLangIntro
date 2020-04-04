package main

import (
	"GoLangIntro/DistributedCrawler/client"
	"GoLangIntro/DistributedCrawler/config"
	"GoLangIntro/SingleThreadCrawler/engine"
	"GoLangIntro/SingleThreadCrawler/parser"
	"GoLangIntro/SingleThreadCrawler/scheduler"
	"fmt"
)

func main() {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(
		engine.Request{
			Url: "http://newcar.xcar.com.cn/",
			Parser: engine.NewFuncParser(
				parser.ParseCarList, "ParseCarList"),
		})
}
