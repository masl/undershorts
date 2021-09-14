package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/masl/undershorts/internal/db"
)

type PostBody struct {
	LongUrl   string `json:"longUrl"`
	ShortPath string `json:"shortPath"`
}

func ShortenEndpoint(router *mux.Router) {
	// POST shorts data
	router.HandleFunc("/shorten", func(rw http.ResponseWriter, r *http.Request) {
		// Set up authorization
		un, pw, ok := r.BasicAuth()
		if !ok {
			fmt.Println("Error parsing basic auth")
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		if un != db.GetEnv("AUTH_USERNAME", "username") {
			fmt.Println("Error parsing basic auth")
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		if pw != db.GetEnv("AUTH_PASSWORD", "password") {
			fmt.Println("Error parsing basic auth")
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		fmt.Println("Shorten POST request sent")
		var latestErr error

		// Read in request data
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

		// Check path existence
		if db.Exist(pb.ShortPath) {
			latestErr = fmt.Errorf("path already exists")
		}

		// Write data to redis
		err = db.SetURL(pb.ShortPath, pb.LongUrl)
		if err != nil {
			fmt.Println("Error while writing redis db:", err)
			latestErr = err
		}

		// Write creation time to redis
		err = db.SetURL(pb.ShortPath+":time", time.Now().String())
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
}
