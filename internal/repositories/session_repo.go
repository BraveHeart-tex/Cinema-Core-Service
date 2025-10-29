package repositories

import (
	"context"
	"time"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"gorm.io/gorm"
)

type SessionRepository struct {
	BaseRepository
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *SessionRepository) CreateSession(ctx context.Context, session *models.Session) (*models.Session, error) {
	if err := r.DB(ctx).Create(session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func (r *SessionRepository) DeleteSession(ctx context.Context, sessionID string) error {
	return r.DB(ctx).Where("id = ?", sessionID).Delete(&models.Session{}).Error
}

func (r *SessionRepository) GetSession(ctx context.Context, sessionID string) (*models.Session, error) {
	var sesion models.Session
	if err := r.DB(ctx).Where("id = ?", sessionID).First(&sesion).Error; err != nil {
		return nil, err
	}
	return &sesion, nil
}

func (r *SessionRepository) UpdateSessionLastVerifiedAt(ctx context.Context, sessionID string) error {
	return r.DB(ctx).Model(&models.Session{}).Where("id = ?", sessionID).Update("last_verified_at", time.Now()).Error
}

func (r *SessionRepository) DeleteExpiredSessions(ctx context.Context, now time.Time) error {
	return r.DB(ctx).Where("expires_at <= ?", now).Delete(&models.Session{}).Error
}
