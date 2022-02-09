package db

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

// New redis client
func New() *redis.Client {
	redisOptions, err := redis.ParseURL(GetEnv("UNDERSHORTS_REDIS_URL", "redis://:PASSWORD@undershorts_redis:6379"))
	if err != nil {
		panic(err)
	}

	log.Println("Starting redis on", redisOptions.Addr)
	return redis.NewClient(redisOptions)
}

var ctx = context.Background()
var RedisClient *redis.Client

// Return creation timestamp by short path
func GetTime(path string) (timestamp time.Time, err error) {
	rawTime, err := RedisClient.Get(ctx, path+":time").Result()
	if err == redis.Nil || err != nil {
		return
	}
	timestamp, err = time.Parse(time.RFC3339, rawTime)
	if err != nil {
		return
	}
	return
}

// Return existence of path
func Exist(path string) (exists bool) {
	code, _ := RedisClient.Exists(ctx, path).Result()
	return code != 0
}

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
		if !strings.HasSuffix(iter.Val(), ":time") {
			allKeys = append(allKeys, iter.Val())
		}
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
