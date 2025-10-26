package app

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services"
	"github.com/gin-gonic/gin"
)

type App struct {
	UserHandler    *handlers.UserHandler
	AdminHandler   *handlers.AdminHandler
	SessionService *services.SessionService
	UserService    *services.UserService
}

func SetupRouter(appCtx *App) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	RegisterAuthRoutes(api, appCtx.UserHandler, appCtx.SessionService)
	RegisterAdminRoutes(api, appCtx.AdminHandler, appCtx.SessionService, appCtx.UserService)

	return r
}
