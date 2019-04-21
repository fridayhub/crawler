package worker

import (
	"github.com/hakits/crawler/engine"
	"log"
)

type CrawlService struct {}

func (CrawlService) Process(req Request, result *ParseResult) error {
	log.Printf("Process start, name:%s", req.Parser.Name)
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		log.Printf("err:", err)
		return err
	}
	log.Printf("engineReq:%v", engineReq)
	parseResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	log.Printf("parseResult:%v", parseResult)
	*result = SerializeResult(parseResult)
	return nil
}
