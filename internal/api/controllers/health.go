package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET health
func GetHealth(ctx *gin.Context) {
	// TODO: Add DB health checks
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
