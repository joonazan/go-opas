// +build heroku

package main

import (
	"encoding/json"
	"gopkg.in/redis.v3"
	"log"
	"os"
)

func lataaEdistymiset() Edistymiset {

	e := Edistymiset{}
	bytes, err := redisClient.Get("all").Result()
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

var redisClient = uusiRedis()

func uusiRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDISCLOUD_ADDRESS"),
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

	return redisClient.Set("all", data, 0).Err()
}
