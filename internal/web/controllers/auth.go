package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostSingup(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "not implemented"})
	return
}

func PostLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "not implemented"})
	return
}
