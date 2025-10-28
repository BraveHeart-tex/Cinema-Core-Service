package showtimes

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/utils"
)

func BuildShowtimeResponse(showtime *models.Showtime) ShowtimeResponse {
	return ShowtimeResponse{
		ID:        showtime.ID,
		MovieID:   showtime.MovieID,
		TheaterID: showtime.TheaterID,
		StartTime: utils.ToRFC3339Ptr(&showtime.StartTime),
		EndTime:   utils.ToRFC3339Ptr(&showtime.EndTime),
		BasePrice: showtime.BasePrice,
		CreatedAt: utils.ToRFC3339Ptr(&showtime.CreatedAt),
	}
}
