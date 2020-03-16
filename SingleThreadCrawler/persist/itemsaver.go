package persist

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 1
		for {
			item := <-out
			log.Printf("Got Items: No. %d, Content %v", itemCount, item)
			itemCount++
		}
	}()
	return out
}
