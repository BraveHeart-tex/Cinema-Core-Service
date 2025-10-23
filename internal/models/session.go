package models

import "time"

type Session struct {
	ID             string    `gorm:"primaryKey"`
	SecretHash     []byte    `gorm:"not null"`
	LastVerifiedAt time.Time `gorm:"autoUpdateTime"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

type SessionWithToken struct {
	Session
	Token string `gorm:"-" json:"token,omitempty"`
}
