package app

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers"
	"github.com/gin-gonic/gin"
)

type App struct {
	UserHandler *handlers.UserHandler
}

func SetupRouter(appCtx *App) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	RegisterUserRoutes(api, appCtx.UserHandler)

	return r
}
