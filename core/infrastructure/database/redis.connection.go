package database

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func RedisConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Panicf("cannot connect to redis: %s", err)
	}
	log.Printf("connect to redis: %s", pong)
	return rdb
}
