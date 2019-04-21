package main

import (
	"fmt"
	"github.com/hakits/crawler/config"
	"github.com/hakits/crawler/rpccall"
	"github.com/hakits/crawler/worker"
	"testing"
	"time"
)

func TestCrawlSercie(t *testing.T)  {
	const host=":9000"
	go rpccall.RpcServer(
		host,
		worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpccall.NewClient(host)
	if err != nil {
		panic(err)
	}
	var args []string
	args = append(args, "Golanggggg")
	args = append(args, "https://www.zhipin.com/job_detail/9f725c45aa83d9751nN83t-8F1E~.html")
	req := worker.Request{
		Url:    "https://www.zhipin.com/job_detail/9f725c45aa83d9751nN83t-8F1E~.html",
		Parser: worker.SerializedParser{
			Name: "ProfileParser",
			Args: args,
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
