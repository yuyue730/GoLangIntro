package engine

import (
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		if isDuplicate(r.Url) {
			log.Printf("Duplicate url: %s", r.Url)
			continue
		}

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		log.Printf("Size of requests = %d", len(parseResult.Requests))
		for _, req := range parseResult.Requests {
			log.Printf("Got new Url to go deep into %s", req.Url)
		}

		for _, item := range parseResult.Items {
			log.Printf("Got new Item %v", item)
		}
	}
}
