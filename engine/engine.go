package engine

import (
	"log"
	"go_ex/job_crawler/fetcher"
)

func Run(seed ...Request) {
	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching:%s", r.Url)
		body, err := fetcher.Fetcher(r.Url)
		if err != nil {
			log.Printf("Fether:error fetch url %s:%v", r.Url, err)
			continue
		}
		parseReulst := r.ParserFunc(body)
		requests = append(requests, parseReulst.Requests...)

		for _, item := range parseReulst.Items {
			log.Printf("Got item: %v", item)
		}
	}
}
