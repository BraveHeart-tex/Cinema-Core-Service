package handlers

import (
	"net/http"
	"strconv"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/audit"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/dto/movies"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/middleware"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/responses"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	service *services.AdminService
}

func NewAdminHandler(service *services.AdminService) *AdminHandler {
	return &AdminHandler{service: service}
}

func (h *AdminHandler) PromoteUser(ctx *gin.Context) {
	ctxData, exists := ctx.Get(middleware.SessionContextKey)
	var performedByID uint
	var performedByEmail string
	if exists && ctxData != nil {
		user := ctxData.(map[string]any)["user"].(*models.User)
		performedByID = user.Id
		performedByEmail = user.Email
	}

	userIDStr := ctx.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		audit.LogEvent(ctx, audit.AuditEvent{
			Event:    "admin.promote_user.failure",
			Success:  false,
			ErrorMsg: "invalid user id",
			Metadata: map[string]any{
				"performedByID":    performedByID,
				"performedByEmail": performedByEmail,
			},
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

func (h *AdminHandler) CreateMovie(ctx *gin.Context) {
	var req movies.CreateMovieRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	ctxData, exists := ctx.Get(middleware.SessionContextKey)
	var performedByID uint
	var performedByEmail string
	if exists && ctxData != nil {
		user := ctxData.(map[string]any)["user"].(*models.User)
		performedByID = user.Id
		performedByEmail = user.Email
	}

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
