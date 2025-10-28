package models

import "time"

type Theater struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"varchar(255);uniqueIndex;not null"`
	CreatedAt time.Time
}
