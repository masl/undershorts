package main

import (
	"log/slog"
	"os"

	"github.com/masl/undershorts/internal/db"
	"github.com/masl/undershorts/internal/web"
	embd "github.com/masl/undershorts/web"
)

func main() {
	// Database client
	postgres, err := db.NewPostgres()
	if err != nil {
		slog.Error("connection to database could not be established", "error", err)
		os.Exit(1)
	}
	defer postgres.Close()

	// Web filesystem
	webFS, err := embd.WebFS()
	if err != nil {
		slog.Error("getting web filesystem failed", "error", err)
		os.Exit(1)
	}

	// Serve http server
	err = web.Serve(postgres, webFS)
	if err != nil {
		slog.Error("starting webserver failed", "error", err)
		os.Exit(1)
	}
}
