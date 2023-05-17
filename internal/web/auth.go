package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (w *WebServer) PostSignup(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "not implemented"})
	return
}

func (w *WebServer) PostLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "not implemented"})
	return
}
