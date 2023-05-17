package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// auth middleware for validating jwt token
func (s *WebServer) Auth(ctx *gin.Context) {
	// get token from cookie
	tokenString, err := ctx.Cookie("token")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// check validity
	// get user by email in claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user, err := s.store.GetUserByEmail(claims["email"].(string))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// set user in context and continue
		ctx.Set("user", user)
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
