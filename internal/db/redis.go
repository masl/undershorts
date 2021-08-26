package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// New redis client
func New() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

var ctx = context.Background()
var RedisClient *redis.Client

// Set a path and it's corresponding long url
func SetURL(client *redis.Client, path string, url string) (err error) {
	err = client.Set(ctx, path, url, 0).Err()
	if err != nil {
		return
	}
	return
}

// Return long url by shorts path
func GetURL(client *redis.Client, path string) (url string, err error) {
	url, err = client.Get(ctx, path).Result()
	if err == redis.Nil || err != nil {
		return
	}
	return
}

func GetAllURLS(client *redis.Client) (allKeys []string, err error) {
	iter := client.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		allKeys = append(allKeys, iter.Val())
	}
	if err = iter.Err(); err != nil {
		return
	}
	return
}
