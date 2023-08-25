package middlewares

import (
	"os"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InterceptApi(ctx *gin.Context) {
	keyName := os.Getenv("API_KEY_NAME")

	validKey := os.Getenv("API_KEY")

	headerKey := ctx.GetHeader(keyName)

	if headerKey != validKey {
		ctx.AbortWithStatus(http.StatusForbidden)
		return 
	}

	ctx.Next()
}