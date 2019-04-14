package persist

import (
	"github.com/go-redis/redis"
	"log"
)

type ItemSaveService struct {
	Client *redis.Client
	Index string
}

func (i *ItemSaveService)Save(item string, result *string)error {
	i.Client.Info()
	//log.Printf("info:%v", info)
	sadd := i.Client.SAdd(i.Index, item)
	log.Printf("sadd:%v", sadd)

	*result = "ok"
	return nil
}
