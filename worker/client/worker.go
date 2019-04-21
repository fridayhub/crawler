package client

import (
	"fmt"
	"github.com/hakits/crawler/config"
	"github.com/hakits/crawler/engine"
	"github.com/hakits/crawler/rpccall"
	"github.com/hakits/crawler/worker"
	"log"
)

func CreateProcessor() (engine.Processor, error) {
	log.Print("CreateProcessor start...")
	client, err := rpccall.NewClient(fmt.Sprintf(":%d", config.Worker0))
	if err != nil {
		return nil, err
	}

	return func(req engine.Request)(engine.ParseResult, error) {
		sReq := worker.SerializedRequest(req)
		var sResult worker.ParseResult

		log.Print("Before call rpc...")
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		log.Printf("get sResult:%v", sResult)

		return  worker.DeserializeResult(sResult), nil
	}, nil
}