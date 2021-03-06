package main

import (
	"context"
	"log"

	"github.com/masl/undershorts/internal/db"
	"github.com/masl/undershorts/internal/web"
)

func main() {
	// Create redis
	db.RedisClient = db.New()
	err := db.RedisClient.Ping(context.Background()).Err()
	if err != nil {
		log.Println("A Problem with the redis connection occurred:", err)
		return
	}

	// Serve http server
	panic(web.Serve())
}
