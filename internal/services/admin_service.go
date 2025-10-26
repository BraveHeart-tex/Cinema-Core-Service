package services

import (
	"errors"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/domainerrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
)

type AdminService struct {
	userRepo *repositories.UserRepository
}

func NewAdminService(userRepo *repositories.UserRepository) *AdminService {
	return &AdminService{userRepo: userRepo}
}

func (s *AdminService) PromoteToAdmin(userID uint) error {
	user, err := s.userRepo.FindById(userID)
	if err != nil {
		if errors.Is(err, domainerrors.ErrNotFound) {
			return NewNotFound("target user not found")
		}
		return NewInternalError("failed to fetch user")
	}
	if user == nil {
		return NewNotFound("target user not found")
	}

	if user.Role == models.AdminRole {
		return NewConflict("target user is already an admin")
	}

	return s.userRepo.UpdateRole(user.Id, models.AdminRole)
}
