package admin

import (
	"net/http"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/audit"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/dto/movies"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/responses"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services"
	"github.com/gin-gonic/gin"
)

type AdminMovieHandler struct {
	service *services.AdminService
	*AdminBaseHandler
}

func NewAdminMovieHandler(service *services.AdminService) *AdminMovieHandler {
	return &AdminMovieHandler{
		service:          service,
		AdminBaseHandler: &AdminBaseHandler{},
	}
}

func (h *AdminMovieHandler) CreateMovie(ctx *gin.Context) {
	var req movies.CreateMovieRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	performedByID, performedByEmail := h.getCurrentAdmin(ctx)

	movie, err := h.service.CreateMovie(req)
	if err != nil {
		if se, ok := err.(*services.ServiceError); ok {
			audit.LogAdminAction(ctx, audit.AdminAuditParams{
				Action:           "create_movie",
				Success:          false,
				PerformedByID:    performedByID,
				PerformedByEmail: performedByEmail,
				ErrorMsg:         se.Message,
			})
			responses.Error(ctx, se.Code, se.Message)
			return
		}
		audit.LogAdminAction(ctx, audit.AdminAuditParams{
			Action:           "create_movie",
			Success:          false,
			PerformedByID:    performedByID,
			PerformedByEmail: performedByEmail,
			ErrorMsg:         err.Error(),
		})
		responses.Error(ctx, http.StatusInternalServerError, "Failed to create movie")
		return
	}

	audit.LogAdminAction(ctx, audit.AdminAuditParams{
		Action:           "create_movie",
		Success:          true,
		PerformedByID:    performedByID,
		PerformedByEmail: performedByEmail,
	})

	responses.Success(ctx, gin.H{
		"message": "movie created successfully",
		"movie":   movies.BuildMovieResponse(movie),
	}, http.StatusCreated)
}
