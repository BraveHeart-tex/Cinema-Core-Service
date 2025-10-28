// Package services provides business logic for managing sessions,
// including creation, validation, and lifecycle management.
package services

import (
	"errors"
	"strings"
	"time"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/utils"
)

type SessionService struct {
	repo *repositories.SessionRepository
}

func NewSessionService(repo *repositories.SessionRepository) *SessionService {
	return &SessionService{repo: repo}
}

var (
	inactivityTimeout     = 10 * 24 * time.Hour // 10 days
	activityCheckInterval = 1 * time.Hour       // 1 hour
)

func (s *SessionService) CreateSession(userID uint) (*models.SessionWithToken, error) {
	now := time.Now()
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
	}

	createdSession, err := s.repo.CreateSession(session)
	if err != nil {
		return nil, err
	}

	return &models.SessionWithToken{
		Session: *createdSession,
		Token:   token,
	}, nil
}

func (s *SessionService) ValidateSessionToken(token string) (*models.Session, error) {
	now := time.Now()

	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return nil, errors.New("invalid token format")
	}
	sessionID, sessionSecret := parts[0], parts[1]

	session, err := s.repo.GetSession(sessionID)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, nil
	}

	// Verify the secret
	tokenSecretHash := utils.HashSecret(sessionSecret)
	if !utils.ConstantTimeEqual(tokenSecretHash, session.SecretHash) {
		return nil, nil
	}

	if time.Since(session.LastVerifiedAt) >= activityCheckInterval {
		session.LastVerifiedAt = now
		_ = s.repo.UpdateSessionLastVerifiedAt(sessionID)
	}

	return session, nil
}

func (s *SessionService) GetSession(sessionID string) (*models.Session, error) {
	session, err := s.repo.GetSession(sessionID)

	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, nil
	}

	if time.Since(session.LastVerifiedAt) >= inactivityTimeout {
		_ = s.repo.DeleteSession(sessionID)
		return nil, nil
	}

	return session, nil
}
