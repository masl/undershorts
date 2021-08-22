package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/masl/undershorts/internal/handler"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/undershorts": "https://github.com/masl/undershorts",
		"/author":      "https://github.com/masl",
	}
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	// Use flag to pass a yaml file
	defaultYAMLPath := "./paths.yaml"
	flagYAML := flag.String("p", defaultYAMLPath, "The location of a YAML file configuration for paths and urls.")
	flag.Parse()

	yamlContent, err := ioutil.ReadFile(*flagYAML)
	if err != nil {
		panic(err)
	}

	yamlHandler, err := handler.YAMLHandler([]byte(yamlContent), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on http://localhost:8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
