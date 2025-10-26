package app

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers/admin"
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

func RegisterAdminRoutes(router *gin.RouterGroup,
	adminHandler *admin.AdminHandler,
	sessionService *services.SessionService,
	userService *services.UserService,
) {
	adminGroup := router.Group("/admin")
	adminGroup.Use(middleware.SessionAuthMiddleware(sessionService, userService), middleware.RoleMiddleware("admin"))

	// All routes are grouped by domain (users, movies, genres, tickets)
	admin.RegisterAdminRoutes(adminGroup, adminHandler)
}
