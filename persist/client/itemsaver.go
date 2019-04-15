package client

import (
	"github.com/hakits/crawler/engine"
	"github.com/hakits/crawler/rpccall"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client , err := rpccall.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver got item #%d: %v", itemCount, item)
			itemCount ++

			//Call RPC to save item
			result := ""
			err = client.Call("ItemSaveService.Save", item, &result)
			if err != nil {
				log.Printf("Call rpc error:%v", err)
			}
			log.Printf(result)
		}
	}()

	return out, nil
}
