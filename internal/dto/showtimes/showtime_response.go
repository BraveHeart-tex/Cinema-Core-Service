package showtimes

type ShowtimeResponse struct {
	ID        uint    `json:"id"`
	MovieID   uint    `json:"movie_id"`
	TheaterID uint    `json:"theater_id"`
	StartTime string  `json:"start_time"`
	EndTime   string  `json:"end_time"`
	BasePrice float64 `json:"base_price"`
	CreatedAt string  `json:"created_at"`
}
