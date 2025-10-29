package repositories

import (
	"fmt"

	dbutils "github.com/BraveHeart-tex/Cinema-Core-Service/internal/dbUtils"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/domainerrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"gorm.io/gorm"
)

type GenreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) *GenreRepository {
	return &GenreRepository{db: db}
}

func (r *GenreRepository) CreateGenre(genre *models.Genre) (*models.Genre, error) {
	err := r.db.Create(genre).Error
	if err != nil {
		if dbutils.IsUniqueConstraintViolationError(err) {
			return nil, fmt.Errorf("genre with name '%s' already exists: %w", genre.Name, domainerrors.ErrConflict)
		}
		return nil, err
	}
	return genre, nil
}

func (r *GenreRepository) FindMany(genreIDs []uint64) ([]models.Genre, error) {
	var genres []models.Genre
	if err := r.db.Find(&genres, genreIDs).Error; err != nil {
		return nil, err
	}
	return genres, nil
}

func (r *GenreRepository) FindById(genreID uint64) (*models.Genre, error) {
	var genre models.Genre
	if err := r.db.First(&genre, genreID).Error; err != nil {
		return nil, err
	}
	return &genre, nil
}

func (r *GenreRepository) Delete(genreID uint64) error {
	return r.db.Delete(&models.Genre{}, genreID).Error
}

func (r *GenreRepository) UpdateGenre(genreID uint64, name string) error {
	result := r.db.Model(&models.Genre{}).
		Where("id = ?", genreID).
		Updates(map[string]any{"name": name})

	if result.Error != nil {
		if dbutils.IsUniqueConstraintViolationError(result.Error) {
			return fmt.Errorf("genre with name '%s' already exists: %w", name, domainerrors.ErrConflict)
		}
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainerrors.ErrNotFound
	}

	return nil
}
