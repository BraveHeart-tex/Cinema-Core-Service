package admin

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services"
	"github.com/gin-gonic/gin"
)

type AdminGenreHandler struct {
	service *services.AdminService
	*AdminBaseHandler
}

func NewAdminGenreHandler(service *services.AdminService) *AdminGenreHandler {
	return &AdminGenreHandler{
		service:          service,
		AdminBaseHandler: &AdminBaseHandler{},
	}
}

func (h *AdminGenreHandler) CreateGenre(ctx *gin.Context) {
	// TODO:
}

func (h *AdminGenreHandler) UpdateGenre(ctx *gin.Context) {
	// TODO:
}

func (h *AdminGenreHandler) DeleteGenre(ctx *gin.Context) {
	// TODO:
}
