package repositories

import (
	"time"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) CreateSession(session *models.Session) (*models.Session, error) {
	if err := r.db.Create(session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func (r *SessionRepository) DeleteSession(sessionId string) error {
	return r.db.Where("id = ?", sessionId).Delete(&models.Session{}).Error
}

func (r *SessionRepository) GetSession(sessionId string) (*models.Session, error) {
	var sesion models.Session
	if err := r.db.Where("id = ?", sessionId).First(&sesion).Error; err != nil {
		return nil, err
	}
	return &sesion, nil
}

func (r *SessionRepository) UpdateSessionLastVerifiedAt(sessionId string) error {
	return r.db.Model(&models.Session{}).Where("id = ?", sessionId).Update("last_verified_at", time.Now()).Error
}
