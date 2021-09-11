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

type PostBody struct {
	LongUrl   string `json:"longUrl"`
	ShortPath string `json:"shortPath"`
}

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
		shortPath := vars["path"]

		all, err := db.GetAllURLS()
		if err != nil {
			fmt.Println("Error while getting keys")
		}

		for k, v := range all {
			if v == shortPath {
				longUrl, err := db.GetURL(shortPath)
				if err != nil {
					continue
				}
				rw.WriteHeader(http.StatusOK)
				json.NewEncoder(rw).Encode(map[string]string{"path": shortPath, "url": longUrl})
				break
			}
			if k >= len(all)-1 {
				rw.WriteHeader(http.StatusNotFound)
				rw.Write([]byte("this path does not exist"))
			}
		}
	}).Methods("GET")

	// POST shorts data
	api.HandleFunc("/shorten", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Shorten POST request sent")
		var latestErr error

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error while getting request body:", err)
			latestErr = err
		}

		pb := new(PostBody)
		err = json.Unmarshal(b, &pb)
		if err != nil {
			fmt.Println("Error while parsing request body:", err)
			latestErr = err
		}

		if db.Exist(pb.ShortPath) {
			latestErr = fmt.Errorf("path already exists")
		}

		err = db.SetURL(pb.ShortPath, pb.LongUrl)
		if err != nil {
			fmt.Println("Error while writing redis db:", err)
			latestErr = err
		}

		if latestErr != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(latestErr.Error()))
		} else {
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(map[string]string{"shorten": "ok"})

			// Register new route
			go func() {
				router.HandleFunc("/"+pb.ShortPath, func(rw http.ResponseWriter, r *http.Request) {
					http.Redirect(rw, r, pb.LongUrl, http.StatusFound)
				})
			}()
		}
	}).Methods("POST")

	// Start http server
	webAddress := db.GetEnv("UNDERSHORTS_WEB_ADDRESS", "0.0.0.0:8000")
	srv := &http.Server{
		Handler: redisHandler,
		Addr:    webAddress,
	}

	fmt.Println("Starting web server on", webAddress)
	return srv.ListenAndServe()
}
