package repositories

import (
	"fmt"

	dbutils "github.com/BraveHeart-tex/Cinema-Core-Service/internal/dbUtils"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/domainerrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"gorm.io/gorm"
)

type TheaterRepository struct {
	db *gorm.DB
}

func NewTheaterRepository(db *gorm.DB) *TheaterRepository {
	return &TheaterRepository{db: db}
}

func (r *TheaterRepository) Create(theater *models.Theater) (*models.Theater, error) {
	err := r.db.Create(theater).Error
	if err != nil {
		if dbutils.IsUniqueConstraintViolationError(err) {
			return nil, fmt.Errorf("theater with name '%s' already exists: %w", theater.Name, domainerrors.ErrConflict)
		}
		return nil, err
	}
	return theater, nil
}

func (r *TheaterRepository) FindAll() ([]models.Theater, error) {
	var theaters []models.Theater
	if err := r.db.Find(&theaters).Error; err != nil {
		return nil, err
	}
	return theaters, nil
}

func (r *TheaterRepository) Delete() error {
	return r.db.Delete(&models.Theater{}).Error
}

func (r *TheaterRepository) Update(id uint, name string) error {
	result := r.db.Model(&models.Theater{}).
		Where("id = ?", id).Updates(map[string]any{"name": name})

	if result.Error != nil {
		if dbutils.IsUniqueConstraintViolationError(result.Error) {
			return fmt.Errorf("theater with name '%s' already exists: %w", name, domainerrors.ErrConflict)
		}
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainerrors.ErrNotFound
	}

	return nil
}
