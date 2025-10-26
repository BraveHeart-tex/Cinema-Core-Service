package repositories

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (r *AdminRepository) PromoteToAdmin(userID uint) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("role", models.AdminRole).Error
}
