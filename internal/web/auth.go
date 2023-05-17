package web

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/masl/undershorts/internal/models"
	"golang.org/x/crypto/bcrypt"
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
	// parse request body
	var userRequest *models.UserRequest
	err := ctx.BindJSON(&userRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get user from db
	user, err := w.store.GetUserByEmail(userRequest.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// compare password hashes
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userRequest.Password))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// set users cookie
	ctx.SetSameSite(http.SameSiteStrictMode)
	ctx.SetCookie("token", tokenString, 3600, "/", os.Getenv("WEB_DOMAIN"), false, true)

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
