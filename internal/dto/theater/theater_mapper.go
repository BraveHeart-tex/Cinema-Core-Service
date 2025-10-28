package theater

import "github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"

func BuildTheaterResponse(theater *models.Theater) TheaterResponse {
	return TheaterResponse{
		ID:   theater.ID,
		Name: theater.Name,
	}
}
