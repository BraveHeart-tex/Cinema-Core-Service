package showtimes

import (
	"time"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/apperrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
)

type Service struct {
	repo        *repositories.ShowtimeRepository
	movieRepo   *repositories.MovieRepository
	theaterRepo *repositories.TheaterRepository
}

// NewService creates a new showtime service with the provided showtime repository, movie repository, and theater repository.
// It panics if any repository is nil to ensure safe construction.
func NewService(repo *repositories.ShowtimeRepository, movieRepo *repositories.MovieRepository, theaterRepo *repositories.TheaterRepository) *Service {
	if repo == nil {
		panic("ShowtimeRepository cannot be nil")
	}
	if movieRepo == nil {
		panic("MovieRepository cannot be nil")
	}
	if theaterRepo == nil {
		panic("TheaterRepository cannot be nil")
	}

	return &Service{
		repo:        repo,
		movieRepo:   movieRepo,
		theaterRepo: theaterRepo,
	}
}

// CreateShowtime creates a new showtime with the provided movieID, theaterID, start time, end time, and base price.
// It returns ServiceError if start time is after end time, base price is negative, movie or theater is not found, or showtime overlaps with another showtime.
// Otherwise, it returns the created showtime object.
func (s *Service) CreateShowtime(movieID, theaterID uint64, start, end time.Time, basePrice float64) (*models.Showtime, error) {
	if start.After(end) || start.Equal(end) {
		return nil, apperrors.NewBadRequest("start time must be before end time")
	}
	if basePrice < 0 {
		return nil, apperrors.NewBadRequest("base price must be non-negative")
	}

	movie, err := s.movieRepo.FindById(movieID)
	if err != nil {
		return nil, apperrors.NewInternalError("failed to fetch movie")
	}
	if movie == nil {
		return nil, apperrors.NewNotFound("movie not found")
	}

	theater, err := s.theaterRepo.FindById(theaterID)
	if err != nil {
		return nil, apperrors.NewInternalError("failed to fetch theater")
	}
	if theater == nil {
		return nil, apperrors.NewNotFound("theater not found")
	}

	exists, err := s.repo.ExistsOverlap(theaterID, start, end)
	if err != nil {
		return nil, apperrors.NewInternalError("failed to check overlapping showtimes")
	}
	if exists {
		return nil, apperrors.NewConflict("showtime overlaps with another showtime")
	}

	showtime := &models.Showtime{
		MovieID:   movieID,
		TheaterID: theaterID,
		StartTime: start,
		EndTime:   end,
		BasePrice: basePrice,
	}

	if err := s.repo.Create(showtime); err != nil {
		return nil, apperrors.NewInternalError("failed to create showtime")
	}

	return showtime, nil
}
