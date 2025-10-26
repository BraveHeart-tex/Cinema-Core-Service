package services

import "github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"

type AdminService struct {
	repo *repositories.AdminRepository
}

func NewAdminService(repo *repositories.AdminRepository) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) PromoteToAdmin(userID uint) error {
	return s.repo.PromoteToAdmin(userID)
}
