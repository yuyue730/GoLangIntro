package persist

import (
	"context"
	"log"

	"github.com/olivere/elastic/v7"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 1
		for {
			item := <-out
			log.Printf("Got Items: No. %d, Content %v", itemCount, item)
			itemCount++

			_, err := save(item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", err, item)
			}
		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false)) // Must turn off sniff in docker mode

	if err != nil {
		return "", err
	}

	resp, err := client.Index().
		Index("car_profile").BodyJson(item).Do(context.Background())

	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
