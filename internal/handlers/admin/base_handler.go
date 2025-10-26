package admin

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/middleware"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/gin-gonic/gin"
)

type AdminBaseHandler struct{}

func (h *AdminBaseHandler) getCurrentAdmin(ctx *gin.Context) (uint, string) {
	val, exists := ctx.Get(middleware.SessionContextKey)
	if exists && val != nil {
		user := val.(map[string]any)["user"].(*models.User)
		return user.Id, user.Email
	}
	return 0, ""
}
