package main

import (
	"GoLangIntro/DistributedCrawler/config"
	"GoLangIntro/DistributedCrawler/persist"
	"GoLangIntro/DistributedCrawler/rpcsupport"
	"flag"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

var port = flag.Int("port", 0, "Port for worker server to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must specify a port")
		return
	}

	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port),
		config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
