package main

import (
	"time"
	"github.com/go-redis/redis"
)

func getClient() (r *redis.Client) {
	r = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return
}

func SetNetPromoterScore(id string) (err error) {
	client := getClient()
	currentTime := time.Now()
	currentTimeFormt := currentTime.Format("2006-01-02T15:04:05.000Z")
	err = client.Set(id, currentTimeFormt, 0).Err()
	return
}

func GetNetPromoterScore(id string) (result string, err error ) {
	client := getClient()
	netPromoterScore, err := client.Get(id).Result()
	if err != nil {
		//fmt.Println("[myredis] houve um erro ao buscar o cliente .", err.Error())
		return
	}
	result = netPromoterScore
	return
}

func DeleteNetPromoterScore(id string) (err error) {
	client := getClient()
	err = client.Del(id).Err()
	return
}
