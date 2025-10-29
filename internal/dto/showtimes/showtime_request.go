package showtimes

type CreateShowtimeRequest struct {
	MovieID   uint64  `json:"movie_id" binding:"required"`
	TheaterID uint64  `json:"theater_id" binding:"required"`
	StartTime string  `json:"start_time" binding:"required"` // RFC3339 string
	EndTime   string  `json:"end_time" binding:"required"`   // RFC3339 string
	BasePrice float64 `json:"base_price" binding:"required,gt=0"`
}
