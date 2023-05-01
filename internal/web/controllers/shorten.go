package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masl/undershorts/internal/db"
)

type PostBody struct {
	LongUrl   string `json:"longUrl" binding:"required,url"`
	ShortPath string `json:"shortPath" binding:"required,alphanum"`
}

// POST shorten
func PostShorten(ctx *gin.Context) {
	var requestBody PostBody

	// TODO: auth middleware

	// Read and validate request body
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Shorten request sent:", requestBody.LongUrl)

	// Check path existence
	if db.Exist(requestBody.ShortPath) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Path already exists"})
		return
	}

	// Write content to database
	if err := db.SetURL(requestBody.ShortPath, requestBody.LongUrl); err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Content written sucessfully
	ctx.Writer.WriteHeader(http.StatusCreated)
}

// func ShortenEndpoint(router *mux.Router, mux *mux.Router) {
// 	// POST shorts data
// 	router.HandleFunc("/shorten", func(rw http.ResponseWriter, r *http.Request) {
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
