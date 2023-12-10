package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)
const (
    authorizationHeader = "Authorization"
    bearerPrefix        = "Bearer "
)


func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader(authorizationHeader)

		
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Unauthorized"})
			ctx.Abort()
			return
		}

		splitToken := strings.Split(tokenString, bearerPrefix)
		if len(splitToken) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Invalid token format"})
			ctx.Abort()
			return
		}

		tokenString = splitToken[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Invalid token: " + err.Error(),
			})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Failed to parse",
			})
			ctx.Abort()
			return
		}
		userId, ok := claims["id"].(float64)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Failed to parse user ID",
			})
			ctx.Abort()
			return
		}
		ctx.Set("id", uint(userId))

		ctx.Next()
	}
}
