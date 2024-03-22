package web

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/masl/undershorts/internal/db"
	"github.com/masl/undershorts/internal/utils"
)

func Serve(postgres db.PostgresClient) (err error) {

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

	// TODO: serve frontend
	// TODO: implement api endpoints
	/*
		// Route API endpoints
		api := router.Group("/api")
		{
			v1 := api.Group("/v1")
			{
				v1.GET("/health", controllers.GetHealth)
				v1.GET("/path/:path", controllers.GetPath)
				v1.POST("/shorten", controllers.PostShorten)
			}
		}
	*/

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
