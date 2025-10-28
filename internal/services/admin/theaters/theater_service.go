package theaters

import (
	"errors"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/apperrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/domainerrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/dto/theater"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
)

type Service struct {
	repo *repositories.TheaterRepository
}

func NewService(repo *repositories.TheaterRepository) *Service {
	if repo == nil {
		panic("TheaterRepository cannot be nil")
	}

	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateTheater(req theater.CreateTheaterRequest) (*models.Theater, error) {
	theater := &models.Theater{
		Name: req.Name,
	}

	result, err := s.repo.Create(theater)
	if err != nil {
		if errors.Is(err, domainerrors.ErrConflict) {
			return nil, apperrors.NewConflict("theater with this name already exists")
		}
		return nil, apperrors.NewInternalError("failed to create theater")
	}

	return result, nil
}
