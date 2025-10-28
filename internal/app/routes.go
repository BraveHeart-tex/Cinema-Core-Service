package app

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers/admin"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/middleware"
	sessionServices "github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/session"
	userServices "github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/user"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup, userHandler *handlers.UserHandler, sessionService *sessionServices.SessionService) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", middleware.GuestOnlyMiddleware(sessionService), userHandler.SignUp)
		auth.POST("/signin", middleware.GuestOnlyMiddleware(sessionService), userHandler.SignIn)
	}
}

// RegisterAdminRoutes registers all admin routes with the router group.
// It uses the SessionAuthMiddleware and RoleMiddleware to validate the session and role for each route.
// All routes are grouped by domain (users, movies, genres, tickets)
func RegisterAdminRoutes(router *gin.RouterGroup,
	adminHandler *admin.AdminHandler,
	sessionService *sessionServices.SessionService,
	userService *userServices.UserService,
) {
	adminGroup := router.Group("/admin")
	adminGroup.Use(middleware.SessionAuthMiddleware(sessionService, userService), middleware.RoleMiddleware("admin"))

	// All routes are grouped by domain (users, movies, genres, tickets)
	admin.RegisterAdminRoutes(adminGroup, adminHandler)
}
