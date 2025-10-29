package models

import "time"

const (
	UserRole  = "user"
	AdminRole = "admin"
)

type User struct {
	Id             uint64 `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	Surname        string
	Email          string `gorm:"uniqueIndex;not null"`
	HashedPassword string `gorm:"not null"`
	Role           string `gorm:"type:varchar(20);default:'regular'"`
	CreatedAt      time.Time
}
