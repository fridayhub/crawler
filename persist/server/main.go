package main

import (
	"github.com/go-redis/redis"
	"github.com/hakits/crawler/persist"
	"github.com/hakits/crawler/rpccall"
)

func main() {
	err := serveRpc("localhost:12345")
	if err != nil {
		panic(err)
	}
}

func serveRpc(host string) error {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	return rpccall.RpcServer(host, &persist.ItemSaveService{
		Client: redisCli,
		Index:"zhipin",
	})

}
