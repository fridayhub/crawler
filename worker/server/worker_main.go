package main

import (
	"fmt"
	"github.com/hakits/crawler/config"
	"github.com/hakits/crawler/rpccall"
	"github.com/hakits/crawler/worker"
	"log"
)

func main() {
	log.Fatal(rpccall.RpcServer(
		fmt.Sprintf(":%d", config.Worker0),
		worker.CrawlService{}))
}
