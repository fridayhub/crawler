package main

import (
	"github.com/hakits/crawler/rpccall"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T)  {
	// start ItemsSaver
	const host  = "127.0.0.1:12345"
	go serveRpc(host)
	time.Sleep(time.Second*2)
	//start iteemCLient
	client , err := rpccall.NewClient(host)
	if err != nil {
		panic(err)
	}
	//call save
	result := ""
	err = client.Call("ItemSaveService.Save", "just test!", &result)
	if err != nil || result != "ok" {
		t.Errorf("result:%s, err:%s", result, err)
	}

}
