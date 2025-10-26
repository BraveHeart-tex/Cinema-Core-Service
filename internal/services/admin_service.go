package services

import (
	"errors"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/domainerrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/dto/movies"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
)

type AdminService struct {
	userRepo  *repositories.UserRepository
	genreRepo *repositories.GenreRepository
	movieRepo *repositories.MovieRepository
}

func NewAdminService(userRepo *repositories.UserRepository, genreRepo *repositories.GenreRepository, movieRepo *repositories.MovieRepository) *AdminService {
	return &AdminService{userRepo: userRepo, genreRepo: genreRepo, movieRepo: movieRepo}
}

func (s *AdminService) PromoteToAdmin(userID uint) error {
	user, err := s.userRepo.FindById(userID)
	if err != nil {
		if errors.Is(err, domainerrors.ErrNotFound) {
			return NewNotFound("target user not found")
		}
		return NewInternalError("failed to fetch user")
	}
	if user == nil {
		return NewNotFound("target user not found")
	}

	if user.Role == models.AdminRole {
		return NewConflict("target user is already an admin")
	}

	return s.userRepo.UpdateRole(user.Id, models.AdminRole)
}

func (s *AdminService) CreateMovie(req movies.CreateMovieRequest) (*models.Movie, error) {
	genres, err := s.genreRepo.FindMany(req.GenreIDs)
	if err != nil {
		return nil, NewInternalError("failed to fetch genres")
	}

	if len(genres) != len(req.GenreIDs) {
		return nil, NewBadRequest("some genre IDs are invalid")
	}

	movie := &models.Movie{
		Title:           req.Title,
		Description:     req.Description,
		PosterImage:     req.PosterImage,
		DurationMinutes: req.DurationMinutes,
		Genres:          genres,
	}

	if err := s.movieRepo.CreateMovie(movie); err != nil {
		return nil, NewInternalError("failed to create movie")
	}

	return movie, nil
}
