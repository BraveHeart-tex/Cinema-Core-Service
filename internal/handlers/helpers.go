package handlers

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/middleware"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services"
	"github.com/gin-gonic/gin"
)

// TODO: Move this to dto package
func buildAuthResponse(result *services.UserWithSession) gin.H {
	return gin.H{
		"user": gin.H{
			"id":    result.User.Id,
			"name":  result.User.Name,
			"email": result.User.Email,
			"role":  result.User.Role,
		},
		"session": gin.H{
			"token": result.Session.Token,
		},
	}
}

func getCurrentAdmin(ctx *gin.Context) (id uint, email string) {
	val, exists := ctx.Get(middleware.SessionContextKey)
	if !exists || val == nil {
		return 0, ""
	}
	user := val.(map[string]any)["user"].(*models.User)
	return user.Id, user.Email
}
