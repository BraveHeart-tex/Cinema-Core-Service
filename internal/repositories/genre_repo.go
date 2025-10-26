package repositories

import (
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
	if err := r.db.Create(genre).Error; err != nil {
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
