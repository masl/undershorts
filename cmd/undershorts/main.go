package main

import (
	"context"
	"fmt"

	"github.com/masl/undershorts/internal/db"
	"github.com/masl/undershorts/internal/web"
)

func main() {
	// Create redis
	db.RedisClient = db.New()
	err := db.RedisClient.Ping(context.Background()).Err()
	if err != nil {
		fmt.Println("Problem with redis connection")
		return
	}

	// Serve http server
	panic(web.Serve())
}
