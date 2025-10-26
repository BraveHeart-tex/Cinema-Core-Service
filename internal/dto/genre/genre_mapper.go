package genre

import "github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"

func BuildUpdateGenreResponse(genre *models.Genre) UpdateGenreResponse {
	return UpdateGenreResponse{
		Id:   genre.ID,
		Name: genre.Name,
	}
}
