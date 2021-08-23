package web

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
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
	yamlContent, err := ioutil.ReadFile(defaultYAMLPath)
	if err != nil {
		return
	}

	yamlHandler, err := handler.YAMLHandler([]byte(yamlContent), mapHandler)
	if err != nil {
		return
	}

	// Start http server
	srv := &http.Server{
		Handler: yamlHandler,
		Addr:    "127.0.0.1:8000",
	}

	fmt.Println("Starting the server on http://127.0.0.1:8000")
	return srv.ListenAndServe()
}
