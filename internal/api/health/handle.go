package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: Add DB health checks
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
