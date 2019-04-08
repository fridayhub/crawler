package engine

import (
	"github.com/hakits/crawler/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (se SimpleEngine) Run(seed ...Request) {
	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item: %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching:%s", r.Url)
	body, err := fetcher.Fetcher(r.Url)
	if err != nil {
		log.Printf("Fether:error fetch url %s:%v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}