package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// New redis client
func New() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "tcp://172.18.0.2:6379",
		Password: "",
		DB:       0,
	})
}

var ctx = context.Background()
var RedisClient *redis.Client

// Set a path and it's corresponding long url
func SetURL(path string, url string) (err error) {
	err = RedisClient.Set(ctx, path, url, 0).Err()
	if err != nil {
		return
	}
	return
}

// Return long url by shorts path
func GetURL(path string) (url string, err error) {
	url, err = RedisClient.Get(ctx, path).Result()
	if err == redis.Nil || err != nil {
		return
	}
	return
}

func GetAllURLS() (allKeys []string, err error) {
	iter := RedisClient.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		allKeys = append(allKeys, iter.Val())
	}
	if err = iter.Err(); err != nil {
		return
	}
	return
}
