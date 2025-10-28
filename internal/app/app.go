package app

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers/admin"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/middleware"
	sessionServices "github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/session"
	userServices "github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/user"
	"github.com/gin-gonic/gin"
)

type App struct {
	UserHandler    *handlers.UserHandler
	AdminHandler   *admin.AdminHandler
	SessionService *sessionServices.SessionService
	UserService    *userServices.UserService
}

func SetupRouter(appCtx *App) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.RequestIDMiddleware())

	api := router.Group("/api")
	RegisterAuthRoutes(api, appCtx.UserHandler, appCtx.SessionService)
	RegisterAdminRoutes(api, appCtx.AdminHandler, appCtx.SessionService, appCtx.UserService)

	return router
}
