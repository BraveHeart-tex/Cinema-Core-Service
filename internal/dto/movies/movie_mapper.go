package movies

import "github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"

func BuildMovieResponse(movie *models.Movie) MovieResponse {
	genres := make([]GenreResponse, len(movie.Genres))
	for i, g := range movie.Genres {
		genres[i] = GenreResponse{
			ID:   g.ID,
			Name: g.Name,
		}
	}

	return MovieResponse{
		ID:              movie.ID,
		Title:           movie.Title,
		Description:     movie.Description,
		PosterImage:     movie.PosterImage,
		DurationMinutes: movie.DurationMinutes,
		Genres:          genres,
	}
}
