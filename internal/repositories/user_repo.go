package repositories

import (
	"context"
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

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := r.DB(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	if err := r.DB(ctx).Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, domainerrors.ErrConflict
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindById(ctx context.Context, userID uint64) (*models.User, error) {
	var user models.User
	err := r.DB(ctx).First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateRole(ctx context.Context, userID uint64, newRole string) error {
	return r.DB(ctx).Model(&models.User{}).Where("id = ?", userID).Update("role", newRole).Error
}
