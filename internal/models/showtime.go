package models

import "time"

type Showtime struct {
	ID        uint      `gorm:"primaryKey"`
	MovieID   uint      `gorm:"not null;index:idx_movie_theater_start,unique"`
	Movie     Movie     `gorm:"foreignKey:MovieID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	TheaterID uint      `gorm:"not null;index:idx_movie_theater_start,unique"`
	Theater   Theater   `gorm:"foreignKey:TheaterID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	StartTime time.Time `gorm:"not null;index:idx_movie_theater_start,unique"`
	EndTime   time.Time `gorm:"not null"`
	BasePrice float64   `gorm:"not null"`
	CreatedAt time.Time
}
