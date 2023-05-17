package main

import (
	"context"
	"log"

	"github.com/masl/undershorts/internal/db"
	"github.com/masl/undershorts/internal/storage"
	"github.com/masl/undershorts/internal/web"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	// establish postgres database connection
	pg, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatalln("A problem with the database connection occured:", err)
		return
	}

	if !pg.Ping() {
		log.Fatalln("A problem occured while trying to ping the database:", err)
		return
	}

	if err := pg.Init(); err != nil {
		log.Fatalf("Failed initializing the database")
		return
	}

	// Create redis
	db.RedisClient = db.New()
	err = db.RedisClient.Ping(context.Background()).Err()
	if err != nil {
		log.Println("A Problem with the redis connection occurred:", err)
		return
	}

	// Serve http server
	web := web.NewWebServer(pg)
	panic(web.Serve())
}
