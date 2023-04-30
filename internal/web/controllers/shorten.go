package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masl/undershorts/internal/db"
)

type PostBody struct {
	LongUrl   string `json:"longUrl"`
	ShortPath string `json:"shortPath"`
}

// POST shorten
func PostShorten(ctx *gin.Context) {
	var requestBody PostBody

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := db.SetURL(requestBody.ShortPath, requestBody.LongUrl); err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Writer.WriteHeader(http.StatusCreated)
}

// func ShortenEndpoint(router *mux.Router, mux *mux.Router) {
// 	// POST shorts data
// 	router.HandleFunc("/shorten", func(rw http.ResponseWriter, r *http.Request) {
// 		// TODO: Set up authorization

// 		/*
// 			un, pw, ok := r.BasicAuth()
// 			if !ok {
// 				log.Println("Error parsing basic auth")
// 				rw.WriteHeader(http.StatusUnauthorized)
// 				return
// 			}

// 			if un != db.GetEnv("AUTH_USERNAME", "username") {
// 				log.Println("Error parsing basic auth")
// 				rw.WriteHeader(http.StatusUnauthorized)
// 				return
// 			}

// 			if pw != db.GetEnv("AUTH_PASSWORD", "password") {
// 				log.Println("Error parsing basic auth")
// 				rw.WriteHeader(http.StatusUnauthorized)
// 				return
// 			}
// 		*/

// 		var latestErr error

// 		// Read in request data
// 		b, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			log.Println("Failed getting request body:", err)
// 			latestErr = err
// 		}

// 		pb := new(PostBody)
// 		err = json.Unmarshal(b, &pb)
// 		if err != nil {
// 			log.Println("Failed parsing request body:", err)
// 			latestErr = err
// 		}

// 		log.Println("Shorten request sent:", pb.ShortPath)

// 		// Check path existence
// 		if db.Exist(pb.ShortPath) {
// 			latestErr = fmt.Errorf("path already exists")
// 		}

// 		// Validate short path
// 		if pb.ShortPath == "" {
// 			latestErr = fmt.Errorf("short path cannot be empty")
// 		}

// 		// Validate long url
// 		if pb.LongUrl == "" {
// 			latestErr = fmt.Errorf("long url cannot be empty")
// 		}

// 		// Validate short path length
// 		if len(pb.ShortPath) > 20 {
// 			latestErr = fmt.Errorf("short path cannot be longer than 20 characters")
// 		}

// 		// Validate long url length
// 		if len(pb.LongUrl) > 1000 {
// 			latestErr = fmt.Errorf("long url cannot be longer than 1000 characters")
// 		}

// 		// Validate long url protocol
// 		if !strings.HasPrefix(pb.LongUrl, "http://") && !strings.HasPrefix(pb.LongUrl, "https://") {
// 			latestErr = fmt.Errorf("long url must start with http:// or https://")
// 		}

// 		// Validate short path characters
// 		for _, c := range pb.ShortPath {
// 			if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')) {
// 				latestErr = fmt.Errorf("short path must only contain alphanumeric characters")
// 			}
// 		}

// 		if latestErr != nil {
// 			rw.WriteHeader(http.StatusInternalServerError)
// 			_, err := rw.Write([]byte(latestErr.Error()))
// 			if err != nil {
// 				log.Println("Failed responding:", err)
// 				return
// 			}
// 			log.Println("Failed linking", pb.ShortPath, "to", pb.LongUrl, ":", latestErr)
// 		} else {
// 			// Write data to redis
// 			err = db.SetURL(pb.ShortPath, pb.LongUrl)
// 			if err != nil {
// 				log.Println("Failed writing to redis DB:", err)
// 				latestErr = err
// 			}

// 			// Write creation time to redis
// 			err = db.SetURL(pb.ShortPath+":time", time.Now().Format(time.RFC3339))
// 			if err != nil {
// 				log.Println("Failed writing to redis DB:", err)
// 				latestErr = err
// 			}

// 			rw.WriteHeader(http.StatusOK)
// 			_ = json.NewEncoder(rw).Encode(map[string]string{"shorten": "ok"})

// 			// Register new route
// 			go func() {
// 				mux.HandleFunc("/"+pb.ShortPath, func(rw http.ResponseWriter, r *http.Request) {
// 					http.Redirect(rw, r, pb.LongUrl, http.StatusFound)
// 				})
// 				log.Println("Successfully linked", pb.ShortPath, "to", pb.LongUrl)
// 			}()
// 		}
// 	}).Methods("POST")
// }
