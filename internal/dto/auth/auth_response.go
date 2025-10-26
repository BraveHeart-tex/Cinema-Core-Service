package auth

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services"
	"github.com/gin-gonic/gin"
)

func BuildAuthResponse(result *services.UserWithSession) gin.H {
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
