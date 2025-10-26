package services

import (
	"errors"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/domainerrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
)

type AdminService struct {
	repo     *repositories.AdminRepository
	userRepo *repositories.UserRepository
}

func NewAdminService(repo *repositories.AdminRepository, userRepo *repositories.UserRepository) *AdminService {
	return &AdminService{repo: repo, userRepo: userRepo}
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

	if user.Role == "admin" {
		return NewConflict("target user is already an admin")
	}

	return s.repo.PromoteToAdmin(user.Id)
}
