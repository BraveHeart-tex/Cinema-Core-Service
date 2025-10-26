package admin

import (
	"github.com/gin-gonic/gin"
)

// RegisterAdminRoutes registers all admin routes with the router group.
// This is called from app/routes.go after handler and middleware are set up.
func RegisterAdminRoutes(router *gin.RouterGroup, handler *AdminHandler) {
	// Each domain has its own route registration function for modularity
	registerUserRoutes(router, handler)
	registerMovieRoutes(router, handler)
	registerGenreRoutes(router, handler)
}

// registerUserRoutes registers user management routes.
// Handlers access services via handler.Services.Users
func registerUserRoutes(router *gin.RouterGroup, handler *AdminHandler) {
	users := router.Group("/users")
	{
		users.PUT("/:userID/promote", handler.PromoteUser)
		users.PUT("/:userID/demote", handler.DemoteUser)
	}
}

// registerMovieRoutes registers movie management routes.
// Handlers access services via handler.Services.Movies
func registerMovieRoutes(router *gin.RouterGroup, handler *AdminHandler) {
	movies := router.Group("/movies")
	{
		movies.POST("/", handler.CreateMovie)
		movies.PUT("/:movieID", handler.UpdateMovie)
		movies.DELETE("/:movieID", handler.DeleteMovie)
	}
}

// registerGenreRoutes registers genre management routes.
// Handlers access services via handler.Services.Genres
func registerGenreRoutes(router *gin.RouterGroup, handler *AdminHandler) {
	genres := router.Group("/genres")
	{
		genres.POST("/", handler.CreateGenre)
		genres.PUT("/:genreID", handler.UpdateGenre)
		genres.DELETE("/:genreID", handler.DeleteGenre)
	}
}
