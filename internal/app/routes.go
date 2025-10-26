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
	userHandler *admin.AdminUserHandler,
	movieHandler *admin.AdminMovieHandler,
	genreHandler *admin.AdminGenreHandler,
	sessionService *services.SessionService,
	userService *services.UserService,
) {
	adminGroup := router.Group("/admin")
	adminGroup.Use(middleware.SessionAuthMiddleware(sessionService, userService), middleware.RoleMiddleware("admin"))

	RegisterUserAdminRoutes(adminGroup, userHandler)
	RegisterMovieAdminRoutes(adminGroup, movieHandler)
	RegisterGenreAdminRoutes(adminGroup, genreHandler)
}

func RegisterUserAdminRoutes(router *gin.RouterGroup, handler *admin.AdminUserHandler) {
	users := router.Group("/users")
	{
		users.PUT("/:userID/promote", handler.PromoteUser)
	}
}

func RegisterMovieAdminRoutes(router *gin.RouterGroup, handler *admin.AdminMovieHandler) {
	movies := router.Group("/movies")
	{
		movies.POST("/", handler.CreateMovie)
	}
}

func RegisterGenreAdminRoutes(router *gin.RouterGroup, handler *admin.AdminGenreHandler) {
	genres := router.Group("/genres")
	{
		genres.POST("/", handler.CreateGenre)
		genres.PUT("/:genreID", handler.UpdateGenre)
		genres.DELETE("/:genreID", handler.DeleteGenre)
	}
}
