package middlewares

import (
	"ecommerce-backend/pkg/common/env"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func EnforceAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "you are not allowed to do this operation",
			})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.JWT_SECRET), nil
		})
		if errors.Is(err, jwt.ErrTokenMalformed) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "you are not allowed to do this operation",
			})
			return
		}
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "your token has been expired, please sign-in again",
			})
			return
		}

		username := token.Claims.(jwt.MapClaims)["username"]
		ctx.Set("username", username)
		ctx.Next()
	}
}
