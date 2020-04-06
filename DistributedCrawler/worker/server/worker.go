package main

import (
	"GoLangIntro/DistributedCrawler/rpcsupport"
	"GoLangIntro/DistributedCrawler/worker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, "Port for worker server to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(
		":%d", *port), worker.CrawlService{}))
}
