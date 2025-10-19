package app

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	users := r.Group("/users")
	{
		users.POST("/signup", userHandler.SignUp)
		users.POST("/login", handlers.Login)
	}
}
