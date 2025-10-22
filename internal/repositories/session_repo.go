package repositories

import (
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
