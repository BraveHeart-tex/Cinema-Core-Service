package repositories

import (
	"errors"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/domainerrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	BaseRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.DB().Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := r.DB().Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, domainerrors.ErrConflict
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindById(userID uint) (*models.User, error) {
	var user models.User
	err := r.DB().First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateRole(userID uint, newRole string) error {
	return r.DB().Model(&models.User{}).Where("id = ?", userID).Update("role", newRole).Error
}
