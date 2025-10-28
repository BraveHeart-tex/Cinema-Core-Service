package admin

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/admin/genres"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/admin/movies"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/admin/theaters"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/admin/users"
)

// Services aggregates all admin domain services
type Services struct {
	Users    *users.Service
	Movies   *movies.Service
	Genres   *genres.Service
	Theaters *theaters.Service
}

// NewServices creates a new Services aggregator with all domain services.
// Dependencies are injected from main.go through the repositories.
func NewServices(
	userRepo *repositories.UserRepository,
	genreRepo *repositories.GenreRepository,
	movieRepo *repositories.MovieRepository,
	theaterRepo *repositories.TheaterRepository,
) *Services {
	if userRepo == nil {
		panic("UserRepository cannot be nil")
	}
	if genreRepo == nil {
		panic("GenreRepository cannot be nil")
	}
	if movieRepo == nil {
		panic("MovieRepository cannot be nil")
	}
	if theaterRepo == nil {
		panic("TheaterRepository cannot be nil")
	}
	return &Services{
		Users:    users.NewService(userRepo),
		Movies:   movies.NewService(movieRepo, genreRepo),
		Genres:   genres.NewService(genreRepo),
		Theaters: theaters.NewService(theaterRepo),
	}
}
