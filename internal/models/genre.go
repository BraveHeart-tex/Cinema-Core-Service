package models

type Genre struct {
	ID     uint64  `gorm:"primaryKey"`
	Name   string  `gorm:"type:varchar(100);uniqueIndex;not null"`
	Movies []Movie `gorm:"many2many:movie_genres"`
}
