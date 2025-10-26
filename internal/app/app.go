package app

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services"
	"github.com/gin-gonic/gin"
)

type App struct {
	UserHandler    *handlers.UserHandler
	SessionService *services.SessionService
}

func SetupRouter(appCtx *App) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	RegisterAuthRoutes(api, appCtx.UserHandler, appCtx.SessionService)

	return r
}
