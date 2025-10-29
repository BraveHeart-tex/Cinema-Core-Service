package movies

import (
	"context"
	"errors"
	"time"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/apperrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/db"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/domainerrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/dto/movies"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
)

// Service handles all movie-related admin operations
type Service struct {
	movieRepo *repositories.MovieRepository
	genreRepo *repositories.GenreRepository
	txManager db.TxManager
}

// NewService creates a new movies service with injected dependencies.
// Panics if any repository is nil to ensure safe construction.
func NewService(movieRepo *repositories.MovieRepository, genreRepo *repositories.GenreRepository, txManager db.TxManager) *Service {
	if movieRepo == nil {
		panic("MovieRepository cannot be nil")
	}
	if genreRepo == nil {
		panic("GenreRepository cannot be nil")
	}
	if txManager == nil {
		panic("TxManager cannot be nil")
	}
	return &Service{
		movieRepo: movieRepo,
		genreRepo: genreRepo,
		txManager: txManager,
	}
}

// CreateMovie creates a new movie with the provided genres.
// Returns ServiceError if genre IDs are invalid or creation fails.
func (s *Service) CreateMovie(ctx context.Context, req movies.CreateMovieRequest) (*models.Movie, error) {
	genres, err := s.genreRepo.FindMany(req.GenreIDs)
	if err != nil {
		return nil, apperrors.NewInternalError("failed to fetch genres")
	}

	if len(genres) != len(req.GenreIDs) {
		return nil, apperrors.NewBadRequest("some genre IDs are invalid")
	}

	movie := &models.Movie{
		Title:           req.Title,
		Description:     req.Description,
		PosterImage:     req.PosterImage,
		DurationMinutes: req.DurationMinutes,
		Genres:          genres,
	}

	if err := s.movieRepo.CreateMovie(ctx, movie); err != nil {
		return nil, apperrors.NewInternalError("failed to create movie")
	}

	return movie, nil
}

// UpdateMovie updates an existing movie's details.
func (s *Service) UpdateMovie(ctx context.Context, movieID uint64, data movies.UpdateMovieRequest) (*models.Movie, error) {
	movie, err := s.movieRepo.FindById(ctx, movieID)
	if err != nil {
		if errors.Is(err, domainerrors.ErrNotFound) {
			return nil, apperrors.NewNotFound("movie not found")
		}
		return nil, apperrors.NewInternalError("failed to fetch movie")
	}

	if movie == nil {
		return nil, apperrors.NewNotFound("movie not found")
	}

	if data.Title != "" {
		movie.Title = data.Title
	}
	if data.Description != "" {
		movie.Description = data.Description
	}
	if data.PosterImage != "" {
		movie.PosterImage = data.PosterImage
	}
	if data.DurationMinutes > 0 {
		movie.DurationMinutes = data.DurationMinutes
	}

	movie.UpdatedAt = time.Now()

	err = s.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		if updateErr := s.movieRepo.Update(ctx, movie); updateErr != nil {
			return apperrors.NewInternalError("failed to update movie")
		}

		if len(data.GenreIDs) > 0 {
			if genreErr := s.movieRepo.UpdateMovieGenres(ctx, movieID, data.GenreIDs); genreErr != nil {
				return apperrors.NewInternalError("failed to update genres")
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return movie, nil
}

// DeleteMovie deletes a movie by ID.
func (s *Service) DeleteMovie(movieID uint) error {
	// TODO: Implement delete logic
	return apperrors.NewInternalError("not implemented")
}
