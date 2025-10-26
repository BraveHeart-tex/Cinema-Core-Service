package admin

import (
	"net/http"
	"strconv"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/apperrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/audit"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/dto/genre"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/responses"
	"github.com/gin-gonic/gin"
)

// CreateGenre creates a new genre.
// This handler uses h.Services.Genres service for the business logic.
func (h *AdminHandler) CreateGenre(ctx *gin.Context) {
	var req genre.CreateGenreRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	genre, err := h.Services.Genres.CreateGenre(req.Name)
	if err != nil {
		if se, ok := err.(*apperrors.ServiceError); ok {
			h.logAdminAction(ctx, audit.AdminAuditParams{
				Action:   "create_genre",
				Success:  false,
				ErrorMsg: se.Message,
			})
			responses.Error(ctx, se.Code, se.Message)
			return
		}

		h.logAdminAction(ctx, audit.AdminAuditParams{
			Action:   "create_genre",
			Success:  false,
			ErrorMsg: err.Error(),
		})
		responses.Error(ctx, http.StatusInternalServerError, "Failed to create genre")
		return
	}

	h.logAdminAction(ctx, audit.AdminAuditParams{
		Action:  "create_genre",
		Success: true,
	})

	responses.Success(ctx, gin.H{
		"message": "genre created successfully",
		"genre":   genre,
	}, http.StatusCreated)
}

// UpdateGenre updates an existing genre's name.
func (h *AdminHandler) UpdateGenre(ctx *gin.Context) {
	genreIDStr := ctx.Param("genreID")
	genreID, err := strconv.ParseUint(genreIDStr, 10, 64)
	if err != nil {
		h.logAdminAction(ctx, audit.AdminAuditParams{
			Success:  false,
			Action:   "update_genre",
			ErrorMsg: "invalid genre id",
		})
		responses.Error(ctx, http.StatusBadRequest, "invalid genre id")
		return
	}

	var req genre.UpdateGenreRequest
	if err = ctx.ShouldBindJSON(&req); err != nil {
		h.logAdminAction(ctx, audit.AdminAuditParams{
			Success:  false,
			Action:   "update_genre",
			ErrorMsg: "invalid input: " + err.Error(),
		})
		responses.Error(ctx, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	updatedGenre, err := h.Services.Genres.UpdateGenre(uint(genreID), req.Name)
	if err != nil {
		if se, ok := err.(*apperrors.ServiceError); ok {
			h.logAdminAction(ctx, audit.AdminAuditParams{
				Action:   "update_genre",
				Success:  false,
				ErrorMsg: se.Message,
			})
			responses.Error(ctx, se.Code, se.Message)
			return
		}

		h.logAdminAction(ctx, audit.AdminAuditParams{
			Action:   "update_genre",
			Success:  false,
			ErrorMsg: err.Error(),
		})
		responses.Error(ctx, http.StatusInternalServerError, "failed to update genre")
		return
	}

	h.logAdminAction(ctx, audit.AdminAuditParams{
		Action:  "update_genre",
		Success: true,
	})

	responses.Success(ctx, gin.H{
		"message": "genre updated successfully",
		"genre":   genre.BuildUpdateGenreResponse(updatedGenre),
	})
}

// DeleteGenre deletes a genre by ID.
func (h *AdminHandler) DeleteGenre(ctx *gin.Context) {
	genreIDStr := ctx.Param("genreID")
	genreID, err := strconv.ParseUint(genreIDStr, 10, 64)
	if err != nil {
		h.logAdminAction(ctx, audit.AdminAuditParams{
			Success:      false,
			Action:       "delete_genre",
			ErrorMsg:     "invalid genre id",
			TargetUserID: 0,
		})
		responses.Error(ctx, http.StatusBadRequest, "invalid genre id")
		return
	}

	err = h.Services.Genres.DeleteGenre(uint(genreID))
	if err != nil {
		if se, ok := err.(*apperrors.ServiceError); ok {
			h.logAdminAction(ctx, audit.AdminAuditParams{
				Success:      false,
				Action:       "delete_genre",
				ErrorMsg:     se.Message,
				TargetUserID: uint(genreID),
			})
			responses.Error(ctx, se.Code, se.Message)
			return
		}

		h.logAdminAction(ctx, audit.AdminAuditParams{
			Success:      false,
			Action:       "delete_genre",
			ErrorMsg:     err.Error(),
			TargetUserID: uint(genreID),
		})
		responses.Error(ctx, http.StatusInternalServerError, "Failed to delete genre")
		return
	}

	h.logAdminAction(ctx, audit.AdminAuditParams{
		Success:      true,
		Action:       "delete_genre",
		TargetUserID: uint(genreID),
	})

	responses.Success(ctx, gin.H{
		"message": "genre deleted successfully",
	}, http.StatusOK)
}
