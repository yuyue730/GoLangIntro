package persist

import (
	"GoLangIntro/SingleThreadCrawler/engine"
	"context"
	"log"

	"github.com/olivere/elastic/v7"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false)) // Must turn off sniff in docker mode

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 1
		for {
			item := <-out
			log.Printf("Got Items: No. %d, Content %v", itemCount, item)
			itemCount++

			if itemCount > 3000 {
				break
			}

			err := Save(client, item, index)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", err, item)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, item engine.Item, index string) error {
	indexService := client.Index().Index(index).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
