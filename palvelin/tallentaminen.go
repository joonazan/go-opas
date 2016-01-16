package main

import (
	"encoding/json"
	"github.com/joonazan/go-opas/Godeps/_workspace/src/gopkg.in/redis.v3"
	"log"
	"os"
)

func lataaEdistymiset() Edistymiset {

	e := Edistymiset{}
	bytes, err := client.Get("all").Result()
	if err == nil {
		err := json.Unmarshal([]byte(bytes), &e.data)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Tila ladattu levyltä!")
	} else {
		log.Println("Tuli virhe. Aloitan tyhjästä.")
		e.data = make(map[string]Edistyminen)
	}

	return e
}

var client = uusiRedis()

func uusiRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDISCLOUD_URL"), //"pub-redis-14130.eu-west-1-1.1.ec2.garantiadata.com:14130",
		Password: os.Getenv("REDISCLOUD_PASSWORD"),
		DB:       0,
	})
}

func (e Edistymiset) Tallenna() error {
	e.RLock()
	data, err := json.Marshal(e)
	if err != nil {
		return err
	}
	e.RUnlock()

	return client.Set("all", data, 0).Err()
}
