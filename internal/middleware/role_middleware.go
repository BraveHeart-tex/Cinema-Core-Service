package middleware

import (
	"net/http"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/responses"
	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		val, exists := ctx.Get(SessionContextKey)
		if !exists || val == nil {
			responses.Error(ctx, http.StatusUnauthorized, "unauthorized")
			ctx.Abort()
			return
		}

		ctxData := val.(map[string]interface{})
		user := ctxData["user"].(*models.User)

		if user.Role != requiredRole {
			responses.Error(ctx, http.StatusUnauthorized, "unauthorized")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
