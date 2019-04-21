package engine

import (
	"github.com/hakits/crawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching:%s", r.Url)
	body, err := fetcher.Fetcher(r.Url)
	if err != nil {
		log.Printf("Fether:error fetch url %s:%v", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body), nil
}
