// Package services provides business logic for managing sessions,
// including creation, validation, and lifecycle management.
package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/db"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/utils"
)

type SessionService struct {
	repo      *repositories.SessionRepository
	txManager db.TxManager
}

func NewSessionService(repo *repositories.SessionRepository, txManager db.TxManager) *SessionService {
	return &SessionService{repo: repo, txManager: txManager}
}

var (
	inactivityTimeout     = 10 * 24 * time.Hour // 10 days
	activityCheckInterval = 1 * time.Hour       // 1 hour
	sessionLifetime       = 30 * 24 * time.Hour // 30 days
)

func (s *SessionService) isExpired(session *models.Session) bool {
	now := time.Now()
	if now.After(session.ExpiresAt) {
		return true
	}
	if now.Sub(session.LastVerifiedAt) >= inactivityTimeout {
		return true
	}
	return false
}

func (s *SessionService) CreateSession(ctx context.Context, userID uint64) (*models.SessionWithToken, error) {
	var err error
	now := time.Now()
	expiry := now.Add(sessionLifetime)
	id, err := utils.GenerateSecureRandomString()
	if err != nil {
		return nil, err
	}

	secret, err := utils.GenerateSecureRandomString()
	if err != nil {
		return nil, err
	}

	secretHash := utils.HashSecret(secret)

	token := id + "." + secret

	session := &models.Session{
		ID:         id,
		SecretHash: secretHash,
		CreatedAt:  now,
		UserID:     userID,
		ExpiresAt:  expiry,
	}

	var result *models.SessionWithToken

	err = s.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var createdSession *models.Session
		createdSession, err = s.repo.CreateSession(ctx, session)
		if err != nil {
			return err
		}

		result = &models.SessionWithToken{
			Session: *createdSession,
			Token:   token,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SessionService) ValidateSessionToken(ctx context.Context, token string) (*models.Session, error) {
	now := time.Now()

	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return nil, errors.New("invalid token format")
	}
	sessionID, sessionSecret := parts[0], parts[1]

	session, err := s.repo.GetSession(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, nil
	}

	if s.isExpired(session) {
		_ = s.repo.DeleteSession(ctx, sessionID)
		return nil, nil
	}

	// Verify the secret
	tokenSecretHash := utils.HashSecret(sessionSecret)
	if !utils.ConstantTimeEqual(tokenSecretHash, session.SecretHash) {
		return nil, nil
	}

	if time.Since(session.LastVerifiedAt) >= activityCheckInterval {
		session.LastVerifiedAt = now
		_ = s.repo.UpdateSessionLastVerifiedAt(ctx, sessionID)
	}

	return session, nil
}

func (s *SessionService) GetSession(ctx context.Context, sessionID string) (*models.Session, error) {
	session, err := s.repo.GetSession(ctx, sessionID)

	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, nil
	}

	if s.isExpired(session) {
		_ = s.repo.DeleteSession(ctx, sessionID)
		return nil, nil
	}

	return session, nil
}

func (s *SessionService) CleanupExpiredSessions(ctx context.Context) error {
	return s.repo.DeleteExpiredSessions(ctx, time.Now())
}
