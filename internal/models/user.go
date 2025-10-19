package models

import "time"

const (
	UserRole  = "user"
	AdminRole = "admin"
)

type User struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	Surname        string
	Email          string `gorm:"uniqueIndex;not null"`
	HashedPassword string `gorm:"not null"`
	Role           string `gorm:"type:enum('admin','regular');default:'regular'"`
	CreatedAt      time.Time
}
