package admin

import (
	"net/http"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/apperrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/audit"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/dto/movies"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/responses"
	"github.com/gin-gonic/gin"
)

// CreateMovie creates a new movie with the provided details.
// This handler uses h.Services.Movies service for the business logic.
func (h *AdminHandler) CreateMovie(ctx *gin.Context) {
	var req movies.CreateMovieRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	performedByID, performedByEmail := h.getCurrentAdmin(ctx)

	movie, err := h.Services.Movies.CreateMovie(req)
	if err != nil {
		if se, ok := err.(*apperrors.ServiceError); ok {
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

// UpdateMovie updates an existing movie's details.
func (h *AdminHandler) UpdateMovie(ctx *gin.Context) {
	// TODO: Implement when UpdateMovie service method is ready
	responses.Error(ctx, http.StatusNotImplemented, "not implemented")
}

// DeleteMovie deletes a movie by ID.
func (h *AdminHandler) DeleteMovie(ctx *gin.Context) {
	// TODO: Implement when DeleteMovie service method is ready
	responses.Error(ctx, http.StatusNotImplemented, "not implemented")
}