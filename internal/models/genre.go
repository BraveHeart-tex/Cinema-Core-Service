package models

type Genre struct {
	ID     uint    `gorm:"primaryKey"`
	Name   string  `gorm:"type:varchar(100);uniqueIndex;not null"`
	Movies []Movie `gorm:"many2many:movie_genres"`
}
