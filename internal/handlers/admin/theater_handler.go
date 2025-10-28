package admin

import (
	"net/http"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/apperrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/audit"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/dto/theater"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/responses"
	"github.com/gin-gonic/gin"
)

func (h *AdminHandler) CreateTheater(ctx *gin.Context) {
	var req theater.CreateTheaterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	createdTheater, err := h.Services.Theaters.CreateTheater(req)
	if err != nil {
		if se, ok := err.(*apperrors.ServiceError); ok {
			h.logAdminAction(ctx, audit.AdminAuditParams{
				Action:   "create_theater",
				Success:  false,
				ErrorMsg: se.Message,
			})
			responses.Error(ctx, se.Code, se.Message)
			return
		}

		h.logAdminAction(ctx, audit.AdminAuditParams{
			Action:   "create_theater",
			Success:  false,
			ErrorMsg: err.Error(),
		})
		responses.Error(ctx, http.StatusInternalServerError, "Failed to create theater")
		return
	}

	h.logAdminAction(ctx, audit.AdminAuditParams{
		Action:  "create_theater",
		Success: true,
	})

	responses.Success(ctx, gin.H{
		"message": "Theater created successfully",
		"theater": theater.BuildTheaterResponse(createdTheater),
	})
}
