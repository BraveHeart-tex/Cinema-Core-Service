package handlers

import (
	"net/http"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/cookies"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services"
	"github.com/gin-gonic/gin"
)

type SignUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) SignUp(ctx *gin.Context) {
	var req SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.CreateUser(services.CreateUserData{
		Name:     req.Name,
		Surname:  req.Surname,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if serviceErr, ok := err.(*services.ServiceError); ok {
			ctx.JSON(serviceErr.Code, gin.H{"error": serviceErr.Message})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
		return
	}

	cookies.SetSessionCookie(ctx, result.Session.Token)
	ctx.JSON(http.StatusCreated, gin.H{
		"user": gin.H{
			"id":    result.User.Id,
			"name":  result.User.Name,
			"email": result.User.Email,
			"role":  result.User.Role,
		},
		"session": gin.H{
			"token": result.Session.Token,
		},
	})
}

func (h *UserHandler) SignIn(ctx *gin.Context) {
}
