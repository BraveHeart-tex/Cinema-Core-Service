package app

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup, userHandler *handlers.UserHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", userHandler.SignUp)
		auth.POST("/signin", userHandler.SignIn)
	}
}
