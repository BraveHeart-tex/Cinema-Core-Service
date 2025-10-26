package genres

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/apperrors"
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
// Returns ServiceError if creation fails.
func (s *Service) CreateGenre(name string) (*models.Genre, error) {
	genre := &models.Genre{Name: name}
	
	result, err := s.genreRepo.CreateGenre(genre)
	if err != nil {
		return nil, apperrors.NewInternalError("failed to create genre")
	}

	return result, nil
}

// UpdateGenre updates an existing genre's name.
func (s *Service) UpdateGenre(genreID uint, newName string) (*models.Genre, error) {
	// TODO: Implement update logic
	return nil, apperrors.NewInternalError("not implemented")
}

// DeleteGenre deletes a genre by ID.
func (s *Service) DeleteGenre(genreID uint) error {
	// TODO: Implement delete logic
	return apperrors.NewInternalError("not implemented")
}
