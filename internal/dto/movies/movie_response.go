package movies

type MovieResponse struct {
	ID              uint64          `json:"id"`
	Title           string          `json:"title"`
	Description     string          `json:"description"`
	PosterImage     string          `json:"poster_image"`
	DurationMinutes int             `json:"duration_minutes"`
	Genres          []GenreResponse `json:"genres"`
}

type GenreResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
