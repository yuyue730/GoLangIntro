package main

import (
	"GoLangIntro/DistributedCrawler/persist"
	"GoLangIntro/DistributedCrawler/rpcsupport"
	"log"

	"github.com/olivere/elastic/v7"
)

func main() {
	log.Fatal(serveRpc(":1234", "car_profile"))
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
