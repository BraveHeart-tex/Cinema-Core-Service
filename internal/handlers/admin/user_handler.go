package admin

import (
	"net/http"
	"strconv"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/apperrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/audit"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/responses"
	"github.com/gin-gonic/gin"
)

// PromoteUser promotes a regular user to admin role.
// This handler uses h.Services.Users service for the business logic.
func (h *AdminHandler) PromoteUser(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		h.logAdminAction(ctx, audit.AdminAuditParams{
			Success:      false,
			Action:       "promote_user",
			ErrorMsg:     "invalid user id",
			TargetUserID: 0,
		})
		responses.Error(ctx, http.StatusBadRequest, "invalid user id")
		return
	}

	err = h.Services.Users.PromoteToAdmin(uint(userID))
	if err != nil {
		if se, ok := err.(*apperrors.ServiceError); ok {
			h.logAdminAction(ctx, audit.AdminAuditParams{
				Success:      false,
				Action:       "promote_user",
				ErrorMsg:     se.Message,
				TargetUserID: uint(userID),
			})
			responses.Error(ctx, se.Code, se.Message)
			return
		}

		h.logAdminAction(ctx, audit.AdminAuditParams{
			Success:      false,
			Action:       "promote_user",
			ErrorMsg:     err.Error(),
			TargetUserID: uint(userID),
		})
		responses.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	h.logAdminAction(ctx, audit.AdminAuditParams{
		Success:      true,
		Action:       "promote_user",
		TargetUserID: uint(userID),
	})
	responses.Success(ctx, gin.H{
		"message": "user promoted to admin",
	})
}

// DemoteUser demotes an admin user to regular user role.
func (h *AdminHandler) DemoteUser(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		h.logAdminAction(ctx, audit.AdminAuditParams{
			Success:      false,
			Action:       "demote_user",
			ErrorMsg:     "invalid user id",
			TargetUserID: 0,
		})
		responses.Error(ctx, http.StatusBadRequest, "invalid user id")
		return
	}

	err = h.Services.Users.DemoteFromAdmin(uint(userID))
	if err != nil {
		if se, ok := err.(*apperrors.ServiceError); ok {
			h.logAdminAction(ctx, audit.AdminAuditParams{
				Success:      false,
				Action:       "demote_user",
				ErrorMsg:     se.Message,
				TargetUserID: uint(userID),
			})
			responses.Error(ctx, se.Code, se.Message)
			return
		}

		h.logAdminAction(ctx, audit.AdminAuditParams{
			Success:      false,
			Action:       "demote_user",
			ErrorMsg:     err.Error(),
			TargetUserID: uint(userID),
		})
		responses.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	h.logAdminAction(ctx, audit.AdminAuditParams{
		Success:      true,
		Action:       "demote_user",
		TargetUserID: uint(userID),
	})
	responses.Success(ctx, gin.H{
		"message": "user demoted to regular user",
	})
}
