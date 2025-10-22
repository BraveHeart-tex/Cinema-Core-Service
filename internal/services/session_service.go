package services

import (
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

func (s *SessionService) CreateSession() (*models.SessionWithToken, error) {
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
