package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masl/undershorts/internal/db"
)

// GET path
func (w *WebServer) GetPath(ctx *gin.Context) {
	path := ctx.Param("path")

	if !db.Exist(path) {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	url, err := db.GetURL(path)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ShortPath": path, "LongUrl": url})
}
