package app

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	RegisterUserRoutes(r)

	return r
}
