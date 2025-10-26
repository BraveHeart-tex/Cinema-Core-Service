package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "requestID"

func RequestIDMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.NewString()
		}

		ctx.Set(RequestIDKey, requestID)

		ctx.Writer.Header().Set("X-Request-ID", requestID)

		ctx.Next()
	}
}
