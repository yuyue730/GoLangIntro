package main

import (
	"GoLangIntro/DistributedCrawler/config"
	"GoLangIntro/DistributedCrawler/rpcsupport"
	"GoLangIntro/DistributedCrawler/worker"
	"fmt"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(
		":%d", config.WorkerPort0), worker.CrawlService{}))
}
