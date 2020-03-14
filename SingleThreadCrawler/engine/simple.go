package engine

import (
	"GoLangIntro/SingleThreadCrawler/fetcher"
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

func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	log.Printf("Fetching url %v", r.Url)
	if err != nil {
		log.Printf("Fetcher error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.Parser.Parse(body, r.Url), nil
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
