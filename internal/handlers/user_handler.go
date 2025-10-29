package handlers

import (
	"net/http"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/apperrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/audit"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/cookies"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/dto/auth"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/responses"
	sessionServices "github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/session"
	services "github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/user"
	"github.com/gin-gonic/gin"
)

type SignUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserHandler struct {
	service        *services.UserService
	sessionService *sessionServices.SessionService
}

func NewUserHandler(service *services.UserService, sessionService *sessionServices.SessionService) *UserHandler {
	return &UserHandler{service: service, sessionService: sessionService}
}

func (h *UserHandler) SignUp(ctx *gin.Context) {
	var req SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		audit.LogEvent(ctx, audit.AuditEvent{
			Event:    "auth.signup.failure",
			Email:    req.Email,
			Success:  false,
			ErrorMsg: err.Error(),
		})
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.SignUp(ctx, services.SignUpData{
		Name:     req.Name,
		Surname:  req.Surname,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		var errMsg string
		if serviceErr, ok := err.(*apperrors.ServiceError); ok {
			errMsg = serviceErr.Message
			responses.Error(ctx, serviceErr.Code, serviceErr.Message)
		} else {
			errMsg = "unexpected error"
			responses.Error(ctx, http.StatusInternalServerError, errMsg)
		}

		audit.LogEvent(ctx, audit.AuditEvent{
			Event:    "auth.signup.failure",
			Email:    req.Email,
			Success:  false,
			ErrorMsg: errMsg,
		})
		return
	}

	audit.LogEvent(ctx, audit.AuditEvent{
		Event:   "auth.signup.success",
		UserID:  result.User.Id,
		Email:   result.User.Email,
		Success: true,
	})
	cookies.SetSessionCookie(ctx, result.Session.Token)
	responses.Success(ctx, auth.BuildAuthResponse(result), http.StatusCreated)
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (h *UserHandler) SignIn(ctx *gin.Context) {
	var req SignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		audit.LogEvent(ctx, audit.AuditEvent{
			Event:    "auth.signin.failure",
			Email:    req.Email,
			Success:  false,
			ErrorMsg: err.Error(),
		})
		responses.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.SignIn(ctx, services.SignInData{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		var errMsg string
		if serviceErr, ok := err.(*apperrors.ServiceError); ok {
			errMsg = serviceErr.Message
			responses.Error(ctx, serviceErr.Code, serviceErr.Message)
		} else {
			errMsg = "unexpected error"
			responses.Error(ctx, http.StatusInternalServerError, errMsg)
		}

		audit.LogEvent(ctx, audit.AuditEvent{
			Event:    "auth.signin.failure",
			Email:    req.Email,
			Success:  false,
			ErrorMsg: errMsg,
		})
		return
	}

	audit.LogEvent(ctx, audit.AuditEvent{
		Event:   "auth.signin.success",
		UserID:  result.User.Id,
		Email:   result.User.Email,
		Success: true,
	})
	cookies.SetSessionCookie(ctx, result.Session.Token)
	responses.Success(ctx, auth.BuildAuthResponse(result), http.StatusOK)
}

func (h *UserHandler) SignOut(ctx *gin.Context) {
	token, err := ctx.Cookie(cookies.SessionCookieName)
	if err != nil || token == "" {
		responses.Error(ctx, http.StatusUnauthorized, "unauthorized")
		return
	}

	err = h.sessionService.DeleteSession(ctx.Request.Context(), token)
	if err != nil {
		audit.LogEvent(ctx, audit.AuditEvent{
			Event:    "auth.signout.failure",
			Success:  false,
			ErrorMsg: err.Error(),
		})
		responses.Error(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	cookies.ClearSessionCookie(ctx)
	responses.Success(ctx, gin.H{"message": "signed out successfully"}, http.StatusOK)
}
