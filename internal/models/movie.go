package models

type Movie struct {
	ID              uint64  `gorm:"primaryKey"`
	Title           string  `gorm:"type:varchar(255);not null"`
	Description     string  `gorm:"type:text"`
	PosterImage     string  `gorm:"type:varchar(512)"`
	DurationMinutes int     `gorm:"not null"`
	Genres          []Genre `gorm:"many2many:movie_genres;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
