package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/masl/undershorts/internal/db"
)

type PostBody struct {
	LongUrl   string `json:"longUrl"`
	ShortPath string `json:"shortPath"`
}

func ShortenEndpoint(router *mux.Router, mux *mux.Router) {
	// POST shorts data
	router.HandleFunc("/shorten", func(rw http.ResponseWriter, r *http.Request) {
		// Set up authorization

		/*
			un, pw, ok := r.BasicAuth()
			if !ok {
				log.Println("Error parsing basic auth")
				rw.WriteHeader(http.StatusUnauthorized)
				return
			}

			if un != db.GetEnv("AUTH_USERNAME", "username") {
				log.Println("Error parsing basic auth")
				rw.WriteHeader(http.StatusUnauthorized)
				return
			}

			if pw != db.GetEnv("AUTH_PASSWORD", "password") {
				log.Println("Error parsing basic auth")
				rw.WriteHeader(http.StatusUnauthorized)
				return
			}
		*/

		var latestErr error

		// Read in request data
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Failed getting request body:", err)
			latestErr = err
		}

		pb := new(PostBody)
		err = json.Unmarshal(b, &pb)
		if err != nil {
			log.Println("Failed parsing request body:", err)
			latestErr = err
		}

		// Check path existence
		if db.Exist(pb.ShortPath) {
			latestErr = fmt.Errorf("path already exists")
		}

		// Format short path
		pb.ShortPath = strings.ReplaceAll(pb.ShortPath, "/", "")
		log.Println("Shorten request sent:", pb.ShortPath)

		// Write data to redis
		err = db.SetURL(pb.ShortPath, pb.LongUrl)
		if err != nil {
			log.Println("Failed writing to redis DB:", err)
			latestErr = err
		}

		// Write creation time to redis
		err = db.SetURL(pb.ShortPath+":time", time.Now().Format(time.RFC3339))
		if err != nil {
			log.Println("Failed writing to redis DB:", err)
			latestErr = err
		}

		if latestErr != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			_, err := rw.Write([]byte(latestErr.Error()))
			if err != nil {
				log.Println("Failed responding:", err)
				return
			}
			log.Println("Failed linking", pb.ShortPath, "to", pb.LongUrl, ":", latestErr)
		} else {
			rw.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(rw).Encode(map[string]string{"shorten": "ok"})

			// Register new route
			go func() {
				mux.HandleFunc("/"+pb.ShortPath, func(rw http.ResponseWriter, r *http.Request) {
					http.Redirect(rw, r, pb.LongUrl, http.StatusFound)
				})
				log.Println("Successfully linked", pb.ShortPath, "to", pb.LongUrl)
			}()
		}
	}).Methods("POST")
}
