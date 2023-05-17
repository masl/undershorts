package web

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
func (w *WebServer) PostShorten(ctx *gin.Context) {
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
