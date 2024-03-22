package main

import (
	"log/slog"
	"os"

	"github.com/masl/undershorts/internal/db"
	"github.com/masl/undershorts/internal/web"
)

func main() {
	// Database client
	postgres, err := db.NewPostgres()
	if err != nil {
		slog.Error("connection to database could not be established", "error", err)
		os.Exit(1)
	}
	defer postgres.Close()

	// Serve http server
	err = web.Serve(*postgres)
	if err != nil {
		slog.Error("starting webserver failed", "error", err)
		os.Exit(1)
	}
}
