package movies

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/apperrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/dto/movies"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
)

// Service handles all movie-related admin operations
type Service struct {
	movieRepo *repositories.MovieRepository
	genreRepo *repositories.GenreRepository
}

// NewService creates a new movies service with injected dependencies.
// Panics if any repository is nil to ensure safe construction.
func NewService(movieRepo *repositories.MovieRepository, genreRepo *repositories.GenreRepository) *Service {
	if movieRepo == nil {
		panic("MovieRepository cannot be nil")
	}
	if genreRepo == nil {
		panic("GenreRepository cannot be nil")
	}
	return &Service{
		movieRepo: movieRepo,
		genreRepo: genreRepo,
	}
}

// CreateMovie creates a new movie with the provided genres.
// Returns ServiceError if genre IDs are invalid or creation fails.
func (s *Service) CreateMovie(req movies.CreateMovieRequest) (*models.Movie, error) {
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

	if err := s.movieRepo.CreateMovie(movie); err != nil {
		return nil, apperrors.NewInternalError("failed to create movie")
	}

	return movie, nil
}

// UpdateMovie updates an existing movie's details.
func (s *Service) UpdateMovie(movieID uint, req movies.UpdateMovieRequest) (*models.Movie, error) {
	// TODO: Implement update logic
	return nil, apperrors.NewInternalError("not implemented")
}

// DeleteMovie deletes a movie by ID.
func (s *Service) DeleteMovie(movieID uint) error {
	// TODO: Implement delete logic
	return apperrors.NewInternalError("not implemented")
}
