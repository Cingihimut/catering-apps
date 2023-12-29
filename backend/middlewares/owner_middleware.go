package middlewares

import (

	"net/http"
	"os"
	"log"

	"github.com/gin-gonic/gin"
)
const (
	accessKeyHeader = "X-Access-Key"
)

func getOwnerAccessKey() string {
	key := os.Getenv("OWNER_ACCESS")
	if key == "" {
		return "Owner"
	}
	return key
}
var Key = getOwnerAccessKey()



func OwnerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessKey := ctx.GetHeader(accessKeyHeader)

		log.Printf("Received Access Key: %s", accessKey)
		log.Printf("Expected Access Key: %s", Key)

		if accessKey == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Unauthorized: Missing Access Key",
			})
			ctx.Abort()
			return
		}

		if accessKey != Key {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Unauthorized: Invalid Access Key",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
