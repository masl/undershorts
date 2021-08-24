package main

import (
	"context"
	"fmt"

	"github.com/masl/undershorts/internal/db"
	"github.com/masl/undershorts/internal/web"
)

func main() {
	// Create redis
	redisClient := db.New()
	err := redisClient.Set(context.Background(), "ping", true, 0).Err()
	if err != nil {
		fmt.Println("Problem with redis connection")
		return
	}

	// Serve http server
	panic(web.Serve())
}
