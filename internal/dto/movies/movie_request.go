package movies

type CreateMovieRequest struct {
	Title           string `json:"title" binding:"required,min=1,max=255"`
	Description     string `json:"description" binding:"omitempty,max=2000"`
	PosterImage     string `json:"poster_image" binding:"omitempty,url"`
	DurationMinutes int    `json:"duration_minutes" binding:"required,gte=1,lte=600"`
	GenreIDs        []uint `json:"genre_ids" binding:"required,min=1,dive,gt=0"`
}
