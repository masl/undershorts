package db

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

// New redis client
func New() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     GetEnv("UNDERSHORTS_REDIS_ADDRESS", "127.0.0.1:6379"),
		Password: GetEnv("UNDERSHORTS_REDIS_PASSWORD", ""),
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

// Return all urls
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

// Get env with fallback if env empty
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
