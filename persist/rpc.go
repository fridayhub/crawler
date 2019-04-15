package persist

import (
	"github.com/go-redis/redis"
	"github.com/hakits/crawler/engine"
	"log"
	"reflect"
)

type ItemSaveService struct {
	Client *redis.Client
	Index  string
}


func StructToMap(obj interface{}) map[string]interface{}{
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	log.Printf("oooooooobj:%v", obj)
	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}


func (i *ItemSaveService) Save(item engine.Item, result *string) error {
	payload, ok := item.Payload.(map[string]interface {})
	log.Printf("ssss:%v, %v, %v", payload, ok, reflect.TypeOf(item.Payload))

	payload["Url"] = item.Url
	i.Client.HMSet(item.Id, payload)

	*result = "ok"

	return nil
}
