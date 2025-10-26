package admin

import (
	"net/http"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/apperrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/audit"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/responses"
	"github.com/gin-gonic/gin"
)

// CreateGenre creates a new genre.
// This handler uses h.Services.Genres service for the business logic.
func (h *AdminHandler) CreateGenre(ctx *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required,min=1,max=100"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	performedByID, performedByEmail := h.getCurrentAdmin(ctx)

	genre, err := h.Services.Genres.CreateGenre(req.Name)
	if err != nil {
		if se, ok := err.(*apperrors.ServiceError); ok {
			audit.LogAdminAction(ctx, audit.AdminAuditParams{
				Action:           "create_genre",
				Success:          false,
				PerformedByID:    performedByID,
				PerformedByEmail: performedByEmail,
				ErrorMsg:         se.Message,
			})
			responses.Error(ctx, se.Code, se.Message)
			return
		}

		audit.LogAdminAction(ctx, audit.AdminAuditParams{
			Action:           "create_genre",
			Success:          false,
			PerformedByID:    performedByID,
			PerformedByEmail: performedByEmail,
			ErrorMsg:         err.Error(),
		})
		responses.Error(ctx, http.StatusInternalServerError, "Failed to create genre")
		return
	}

	audit.LogAdminAction(ctx, audit.AdminAuditParams{
		Action:           "create_genre",
		Success:          true,
		PerformedByID:    performedByID,
		PerformedByEmail: performedByEmail,
	})

	responses.Success(ctx, gin.H{
		"message": "genre created successfully",
		"genre":   genre,
	}, http.StatusCreated)
}

// UpdateGenre updates an existing genre's name.
func (h *AdminHandler) UpdateGenre(ctx *gin.Context) {
	// TODO: Implement when UpdateGenre service method is ready
	responses.Error(ctx, http.StatusNotImplemented, "not implemented")
}

// DeleteGenre deletes a genre by ID.
func (h *AdminHandler) DeleteGenre(ctx *gin.Context) {
	// TODO: Implement when DeleteGenre service method is ready
	responses.Error(ctx, http.StatusNotImplemented, "not implemented")
}