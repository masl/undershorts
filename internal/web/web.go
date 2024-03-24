package web

import (
	"io/fs"
	"log/slog"
	"net/http"
	"time"

	"github.com/masl/undershorts/internal/db"
	"github.com/masl/undershorts/internal/utils"
	"github.com/masl/undershorts/internal/web/health"
	"github.com/masl/undershorts/internal/web/short"
	"github.com/masl/undershorts/internal/web/shorten"
)

func Serve(postgres *db.PostgresClient, webFS fs.FS) (err error) {
	// webserver router
	router := http.NewServeMux()

	// handle short url
	router.HandleFunc("GET /{shortURL}", func(w http.ResponseWriter, r *http.Request) {
		shortURL := r.PathValue("shortURL")

		short, err := postgres.GetShortByShortURL(shortURL)
		if err != nil {
			// TODO: handle different error types
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, short.LongURL, http.StatusFound)
	})

	// serve frontend
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, webFS, "index.html")
	})

	router.HandleFunc("GET /assets/{file}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, webFS, "assets/"+r.PathValue("file"))
	})

	// serve api endpoints
	router.HandleFunc("GET /api/v1/health", health.Handle())
	router.HandleFunc("GET /api/v1/{shortURL}", short.Handle(postgres))
	router.HandleFunc("POST /api/v1/shorten", shorten.Handle(postgres))

	// serve webserver
	addr := utils.GetEnv("WEB_ADDRESS", ":8080")
	slog.Info("starting webserver", "address", addr)
	s := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.ListenAndServe()
}
