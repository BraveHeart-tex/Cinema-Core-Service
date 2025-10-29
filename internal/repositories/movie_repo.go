package repositories

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) CreateMovie(movie *models.Movie) error {
	return r.db.Create(movie).Error
}

func (r *MovieRepository) FindById(movieID uint64) (*models.Movie, error) {
	var movie models.Movie
	if err := r.db.First(&movie, movieID).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}
