package app

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("/signup", handlers.SignUp)
		users.POST("/login", handlers.Login)
	}
}
