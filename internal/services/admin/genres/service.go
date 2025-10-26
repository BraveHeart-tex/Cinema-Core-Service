package genres

import (
	"errors"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/apperrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/domainerrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
)

// Service handles all genre-related admin operations
type Service struct {
	genreRepo *repositories.GenreRepository
}

// NewService creates a new genres service with injected dependencies.
// Panics if genreRepo is nil to ensure safe construction.
func NewService(genreRepo *repositories.GenreRepository) *Service {
	if genreRepo == nil {
		panic("GenreRepository cannot be nil")
	}
	return &Service{
		genreRepo: genreRepo,
	}
}

// CreateGenre creates a new genre.
// Returns ServiceError if creation fails, including Conflict error for duplicate names.
func (s *Service) CreateGenre(name string) (*models.Genre, error) {
	genre := &models.Genre{Name: name}
	result, err := s.genreRepo.CreateGenre(genre)
	if err != nil {
		if errors.Is(err, domainerrors.ErrConflict) {
			return nil, apperrors.NewConflict("genre with this name already exists")
		}
		return nil, apperrors.NewInternalError("failed to create genre")
	}

	return result, nil
}

// UpdateGenre updates an existing genre's name.
// Returns ServiceError if genre not found, name conflict, or update fails.
func (s *Service) UpdateGenre(genreID uint, newName string) (*models.Genre, error) {
	if len(newName) == 0 || len(newName) > 100 {
		return nil, apperrors.NewBadRequest("genre name must be between 1 and 100 characters")
	}

	result, err := s.genreRepo.UpdateGenre(genreID, newName)
	if err != nil {

		if errors.Is(err, domainerrors.ErrNotFound) {
			return nil, apperrors.NewNotFound("genre not found")
		}

		if errors.Is(err, domainerrors.ErrConflict) {
			return nil, apperrors.NewConflict("genre with this name already exists")
		}

		return nil, apperrors.NewInternalError("failed to update genre")
	}

	return result, nil
}

// DeleteGenre deletes a genre by ID.
func (s *Service) DeleteGenre(genreID uint) error {
	genre, err := s.genreRepo.FindById(genreID)

	if genre == nil {
		return apperrors.NewNotFound("genre not found")
	}

	if err != nil {
		return apperrors.NewInternalError("failed to delete genre")
	}

	return s.genreRepo.Delete(genre.ID)
}
