package admin

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/audit"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/middleware"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	adminServices "github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/admin"
	"github.com/gin-gonic/gin"
)

// AdminHandler is the main handler that aggregates all admin operations.
// It holds a reference to Services which contains all domain-specific services.
type AdminHandler struct {
	Services *adminServices.Services
}

// NewAdminHandler creates a new AdminHandler with injected services.
// Services should be created in main.go using admin.NewServices(...)
func NewAdminHandler(services *adminServices.Services) *AdminHandler {
	return &AdminHandler{
		Services: services,
	}
}

// getCurrentAdmin extracts the currently authenticated admin user from the context.
// Returns userID and email. This is a shared helper used by all domain handlers.
func (h *AdminHandler) getCurrentAdmin(ctx *gin.Context) (uint, string) {
	val, exists := ctx.Get(middleware.SessionContextKey)
	if exists && val != nil {
		user := val.(map[string]any)["user"].(*models.User)
		return user.Id, user.Email
	}
	return 0, ""
}

// logAdminAction wraps audit.LogAdminAction and automatically fills PerformedByID and PerformedByEmail
// from the current admin context. This reduces duplication across all domain handlers.
// All admin actions should use this method instead of calling audit.LogAdminAction directly.
func (h *AdminHandler) logAdminAction(ctx *gin.Context, params audit.AdminAuditParams) {
	performedByID, performedByEmail := h.getCurrentAdmin(ctx)
	params.PerformedByID = performedByID
	params.PerformedByEmail = performedByEmail
	audit.LogAdminAction(ctx, params)
}

