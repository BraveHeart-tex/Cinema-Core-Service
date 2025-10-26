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

func (r *GenreRepository) FindMany(genreIDs []uint) ([]models.Genre, error) {
	var genres []models.Genre
	if err := r.db.Find(&genres, genreIDs).Error; err != nil {
		return nil, err
	}
	return genres, nil
}

func (r *GenreRepository) FindById(genreID uint) (*models.Genre, error) {
	var genre models.Genre
	if err := r.db.First(&genre, genreID).Error; err != nil {
		return nil, err
	}
	return &genre, nil
}

func (r *GenreRepository) Delete(genreID uint) error {
	return r.db.Delete(&models.Genre{}, genreID).Error
}
