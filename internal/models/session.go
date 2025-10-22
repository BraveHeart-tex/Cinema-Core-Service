package models

import "time"

type Session struct {
	ID         string    `gorm:"primaryKey"`
	SecretHash []byte    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

type SessionWithToken struct {
	Session
	Token string `gorm:"-" json:"token,omitempty"`
}
