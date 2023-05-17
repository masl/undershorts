package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masl/undershorts/internal/models"
)

// api route for signing up with email and password
func (w *WebServer) PostSignup(ctx *gin.Context) {
	// parse request data
	var UserRequest *models.UserRequest
	err := ctx.BindJSON(&UserRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// write user to database
	user, err := w.store.CreateUser(UserRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return data of created user
	ctx.JSON(http.StatusCreated, user)
}

func (w *WebServer) PostLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "not implemented"})
	return
}
