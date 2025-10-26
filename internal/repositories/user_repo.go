package repositories

import (
	"errors"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/domainerrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, domainerrors.ErrConflict
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindById(userID uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}
