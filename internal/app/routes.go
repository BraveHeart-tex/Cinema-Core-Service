package app

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/middleware"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup, userHandler *handlers.UserHandler, sessionService *services.SessionService) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", middleware.GuestOnlyMiddleware(sessionService), userHandler.SignUp)
		auth.POST("/signin", middleware.GuestOnlyMiddleware(sessionService), userHandler.SignIn)
	}
}

func RegisterAdminRoutes(router *gin.RouterGroup, adminHandler *handlers.AdminHandler, sessionService *services.SessionService, userService *services.UserService) {
	admin := router.Group("/admin")
	admin.Use(middleware.SessionAuthMiddleware(sessionService, userService), middleware.RoleMiddleware("admin"))

	users := admin.Group("/users")
	{
		users.PUT("/:userID/promote", adminHandler.PromoteUser)
	}
}
