package admin

import (
	"net/http"
	"strconv"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/audit"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/responses"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services"
	"github.com/gin-gonic/gin"
)

type AdminUserHandler struct {
	service *services.AdminService
	*AdminBaseHandler
}

func NewAdminUserHandler(service *services.AdminService) *AdminUserHandler {
	return &AdminUserHandler{
		service:          service,
		AdminBaseHandler: &AdminBaseHandler{},
	}
}

func (h *AdminUserHandler) PromoteUser(ctx *gin.Context) {
	performedByID, performedByEmail := h.getCurrentAdmin(ctx)

	userIDStr := ctx.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		audit.LogAdminAction(ctx, audit.AdminAuditParams{
			Success:          false,
			Action:           "promote_user",
			ErrorMsg:         "invalid user id",
			TargetUserID:     0,
			PerformedByID:    performedByID,
			PerformedByEmail: performedByEmail,
		})
		responses.Error(ctx, http.StatusBadRequest, "invalid user id")
		return
	}

	err = h.service.PromoteToAdmin(uint(userID))
	if err != nil {
		if se, ok := err.(*services.ServiceError); ok {
			audit.LogAdminAction(ctx, audit.AdminAuditParams{
				Success:          false,
				Action:           "promote_user",
				ErrorMsg:         se.Message,
				TargetUserID:     uint(userID),
				PerformedByID:    performedByID,
				PerformedByEmail: performedByEmail,
			})
			responses.Error(ctx, se.Code, se.Message)
			return
		}

		audit.LogAdminAction(ctx, audit.AdminAuditParams{
			Success:          false,
			Action:           "promote_user",
			ErrorMsg:         err.Error(),
			TargetUserID:     uint(userID),
			PerformedByID:    performedByID,
			PerformedByEmail: performedByEmail,
		})
		responses.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	audit.LogAdminAction(ctx, audit.AdminAuditParams{
		Success:          true,
		Action:           "promote_user",
		TargetUserID:     uint(userID),
		PerformedByID:    performedByID,
		PerformedByEmail: performedByEmail,
	})
	responses.Success(ctx, gin.H{
		"message": "user promoted to admin",
	})
}
