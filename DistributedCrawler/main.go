package main

import (
	itemsaver "GoLangIntro/DistributedCrawler/persist/client"
	"GoLangIntro/DistributedCrawler/rpcsupport"
	worker "GoLangIntro/DistributedCrawler/worker/client"
	"GoLangIntro/SingleThreadCrawler/engine"
	"GoLangIntro/SingleThreadCrawler/parser"
	"GoLangIntro/SingleThreadCrawler/scheduler"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "ItemSaver host")
	workerHosts   = flag.String("worker_hosts", "",
		"Worker hosts (separated by comma)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool)
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(
		engine.Request{
			Url: "http://newcar.xcar.com.cn/",
			Parser: engine.NewFuncParser(
				parser.ParseCarList, "ParseCarList"),
		})
}

// Create a pool of clients continuously fed into a channel of rpc.Client pointer.
// The Channel will be monitored and listened in worker/client/worker
func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
