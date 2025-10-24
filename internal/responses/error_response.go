package responses

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Error(ctx *gin.Context, code int, message string, debug ...string) {
	payload := gin.H{
		"success": false,
		"error":   message,
	}
	if len(debug) > 0 && os.Getenv("APP_ENV") == "development" {
		payload["debug"] = debug[0]
	}
	ctx.JSON(code, payload)
}
