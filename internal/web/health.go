package web

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masl/undershorts/internal/db"
)

type status map[string]bool

// GET health
func (w *WebServer) GetHealth(ctx *gin.Context) {
	status := status{
		"redis":   db.RedisClient.Ping(context.Background()).Err() == nil,
		"storage": w.store.Ping(),
	}

	ctx.JSON(http.StatusOK, status)
}
