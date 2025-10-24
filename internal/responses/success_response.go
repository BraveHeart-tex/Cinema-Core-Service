package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data gin.H, status ...int) {
	code := http.StatusOK
	if len(status) > 0 {
		code = status[0]
	}
	ctx.JSON(code, gin.H{"data": data, "success": true})
}
