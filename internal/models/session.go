package models

import "time"

type Session struct {
	ID             string    `gorm:"primaryKey"`
	UserID         uint64    `gorm:"not null;index"`
	User           User      `gorm:"constraint:OnDelete:CASCADE;"`
	SecretHash     []byte    `gorm:"not null"`
	LastVerifiedAt time.Time `gorm:"autoUpdateTime;index"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

type SessionWithToken struct {
	Session
	Token string `gorm:"-" json:"token,omitempty"`
}
