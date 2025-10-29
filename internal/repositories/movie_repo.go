package repositories

import (
	"context"
	"errors"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/domainerrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"gorm.io/gorm"
)

type MovieRepository struct {
	BaseRepository
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *MovieRepository) CreateMovie(ctx context.Context, movie *models.Movie) error {
	return r.DB(ctx).Create(movie).Error
}

func (r *MovieRepository) FindById(ctx context.Context, movieID uint64) (*models.Movie, error) {
	var movie models.Movie
	if err := r.DB(ctx).First(&movie, movieID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}
		return nil, err
	}
	return &movie, nil
}

func (r *MovieRepository) Update(ctx context.Context, movie *models.Movie) error {
	return r.DB(ctx).Save(movie).Error
}

func (r *MovieRepository) UpdateMovieGenres(ctx context.Context, movieID uint64, genreIDs []uint64) error {
	genres := make([]models.Genre, len(genreIDs))
	for i, genreId := range genreIDs {
		genres[i] = models.Genre{ID: genreId}
	}
	return r.DB(ctx).Model(&models.Movie{ID: movieID}).Association("Genres").Replace(genres)
}
