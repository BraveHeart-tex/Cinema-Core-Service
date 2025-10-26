package handlers

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/middleware"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/gin-gonic/gin"
)

func getCurrentAdmin(ctx *gin.Context) (id uint, email string) {
	val, exists := ctx.Get(middleware.SessionContextKey)
	if !exists || val == nil {
		return 0, ""
	}
	user := val.(map[string]any)["user"].(*models.User)
	return user.Id, user.Email
}
