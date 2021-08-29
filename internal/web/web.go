package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/masl/undershorts/internal/db"
	"github.com/masl/undershorts/internal/handler"
)

func Serve() (err error) {
	router := mux.NewRouter()
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello World!"))
	})

	// Map handler
	pathsToUrls := map[string]string{
		"/undershorts": "https://github.com/masl/undershorts",
		"/author":      "https://github.com/masl",
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
	api := router.PathPrefix("/api").Subrouter()
	// GET status
	api.HandleFunc("/status", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(map[string]bool{"ok": true})
	}).Methods("GET")

	// GET shorts data
	api.HandleFunc("/{path}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars["path"] == "" /* exists */ {
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(map[string]string{"path": vars["path"], "url": ""})
		} else {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte("404 page not found"))
		}
	}).Methods("GET")

	// Start http server
	webAddress := db.GetEnv("UNDERSHORTS_WEB_ADDRESS", "0.0.0.0:8000")
	srv := &http.Server{
		Handler: redisHandler,
		Addr:    webAddress,
	}

	fmt.Println("Starting web server on", webAddress)
	return srv.ListenAndServe()
}
