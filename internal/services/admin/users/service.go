package users

import (
	"errors"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/apperrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/domainerrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
)

// Service handles all user-related admin operations
type Service struct {
	userRepo *repositories.UserRepository
}

// NewService creates a new users service with injected dependencies.
// Panics if userRepo is nil to ensure safe construction.
func NewService(userRepo *repositories.UserRepository) *Service {
	if userRepo == nil {
		panic("UserRepository cannot be nil")
	}
	return &Service{
		userRepo: userRepo,
	}
}

// PromoteToAdmin promotes a regular user to admin role.
// Returns ServiceError if user not found, already admin, or update fails.
func (s *Service) PromoteToAdmin(userID uint) error {
	user, err := s.userRepo.FindById(userID)
	if err != nil {
		if errors.Is(err, domainerrors.ErrNotFound) {
			return apperrors.NewNotFound("target user not found")
		}
		return apperrors.NewInternalError("failed to fetch user")
	}
	if user == nil {
		return apperrors.NewNotFound("target user not found")
	}

	if user.Role == models.AdminRole {
		return apperrors.NewConflict("target user is already an admin")
	}

	if err := s.userRepo.UpdateRole(user.Id, models.AdminRole); err != nil {
		return apperrors.NewInternalError("failed to update user role")
	}

	return nil
}

// DemoteFromAdmin demotes an admin user to regular user role.
func (s *Service) DemoteFromAdmin(userID uint) error {
	user, err := s.userRepo.FindById(userID)
	if err != nil {
		if errors.Is(err, domainerrors.ErrNotFound) {
			return apperrors.NewNotFound("target user not found")
		}
		return apperrors.NewInternalError("failed to fetch user")
	}
	if user == nil {
		return apperrors.NewNotFound("target user not found")
	}

	if user.Role == models.UserRole {
		return apperrors.NewConflict("target user is already a regular user")
	}

	if err := s.userRepo.UpdateRole(user.Id, models.UserRole); err != nil {
		return apperrors.NewInternalError("failed to update user role")
	}

	return nil
}

// GetUserByID fetches a user by their ID.
func (s *Service) GetUserByID(userID uint) (*models.User, error) {
	user, err := s.userRepo.FindById(userID)
	if err != nil {
		if errors.Is(err, domainerrors.ErrNotFound) {
			return nil, apperrors.NewNotFound("user not found")
		}
		return nil, apperrors.NewInternalError("failed to fetch user")
	}
	return user, nil
}
