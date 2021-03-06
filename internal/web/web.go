package web

import (
	"github.com/gorilla/mux"
	"github.com/masl/undershorts/internal/db"
	"github.com/masl/undershorts/internal/handler"
	"github.com/masl/undershorts/internal/web/api"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Serve() (err error) {
	router := mux.NewRouter()

	// Frontend handler
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets"))))

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "./web/index.html")
	})

	// Map handler
	pathsToUrls := map[string]string{
		"undershorts": "https://github.com/masl/undershorts",
		"author":      "https://github.com/masl",
	}

	mapHandler := handler.MapHandler(pathsToUrls, router)

	// YAML handler
	defaultYAMLPath := "./paths.yaml"

	var yamlContent []byte
	var yamlHandler http.Handler

	// Check existence of YAML file
	fileinfo, err := os.Stat(defaultYAMLPath)
	if os.IsNotExist(err) || fileinfo.IsDir() {
		yamlHandler, err = handler.YAMLHandler(make([]byte, 0), mapHandler)
		if err != nil {
			return
		}
	} else {
		yamlContent, err = ioutil.ReadFile(defaultYAMLPath)
		if err != nil {
			return
		}

		yamlHandler, err = handler.YAMLHandler([]byte(yamlContent), mapHandler)
		if err != nil {
			return
		}
	}

	// Redis handler
	redisContent, err := db.GetAllURLS()
	if err != nil {
		return
	}

	redisHandler, err := handler.RedisHandler(redisContent, yamlHandler)
	if err != nil {
		return
	}

	// API handler
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Register API Endpoints
	api.StatusEndpoint(apiRouter)
	api.PathEndpoint(apiRouter)
	api.ShortenEndpoint(apiRouter, router)

	// Start http server
	webAddress := db.GetEnv("UNDERSHORTS_WEB_ADDRESS", "0.0.0.0:8000")
	srv := &http.Server{
		Handler: redisHandler,
		Addr:    webAddress,
	}

	log.Println("Starting web server on", webAddress)
	return srv.ListenAndServe()
}
