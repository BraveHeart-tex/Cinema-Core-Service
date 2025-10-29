package services

import (
	"context"
	"errors"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/apperrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/db"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/domainerrors"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
	services "github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repo           *repositories.UserRepository
	sessionService *services.SessionService
	txManager      db.TxManager
}

func NewUserService(repo *repositories.UserRepository, sessionService *services.SessionService, txManager db.TxManager) *UserService {
	return &UserService{
		repo:           repo,
		sessionService: sessionService,
		txManager:      txManager,
	}
}

type SignUpData struct {
	Name     string
	Surname  string
	Email    string
	Password string
}

type UserWithSession struct {
	User    *models.User
	Session *models.SessionWithToken
}

func (s *UserService) SignUp(ctx context.Context, data SignUpData) (*UserWithSession, error) {
	var err error
	existing, err := s.repo.FindByEmail(data.Email)
	if err != nil && !errors.Is(err, domainerrors.ErrNotFound) {
		return nil, apperrors.NewInternalError("failed to check existing user")
	}

	if existing != nil {
		return nil, apperrors.NewConflict("user already exists with the given email")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperrors.NewInternalError("failed to hash password")
	}

	user := &models.User{
		Name:           data.Name,
		Surname:        data.Surname,
		Email:          data.Email,
		HashedPassword: string(hashed),
		Role:           models.UserRole,
	}

	var result *UserWithSession

	err = s.txManager.WithTransaction(ctx, func(tx *gorm.DB) error {
		txRepo := *s.repo
		txRepo.WithTx(tx)

		var createdUser *models.User
		createdUser, err = txRepo.Create(user)
		if err != nil {
			if errors.Is(err, domainerrors.ErrConflict) {
				return apperrors.NewConflict("user already exists with the given email")
			}
			return apperrors.NewInternalError("failed to create user")
		}

		var session *models.SessionWithToken
		session, err = s.sessionService.CreateSession(ctx, createdUser.Id)
		if err != nil {
			return apperrors.NewInternalError("failed to create session")
		}

		result = &UserWithSession{
			User:    createdUser,
			Session: session,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

type SignInData struct {
	Email    string
	Password string
}

func (s *UserService) SignIn(ctx context.Context, data SignInData) (*UserWithSession, error) {
	var err error

	var user *models.User
	user, err = s.repo.FindByEmail(data.Email)
	if err != nil {
		if errors.Is(err, domainerrors.ErrNotFound) {
			return nil, apperrors.NewUnauthorized("invalid email or password")
		}
		return nil, apperrors.NewInternalError("failed to fetch user")
	}

	if user == nil {
		return nil, apperrors.NewUnauthorized("invalid email or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(data.Password)) != nil {
		return nil, apperrors.NewUnauthorized("invalid email or password")
	}

	var result *UserWithSession

	err = s.txManager.WithTransaction(ctx, func(tx *gorm.DB) error {
		var session *models.SessionWithToken
		session, err = s.sessionService.CreateSession(ctx, user.Id)
		if err != nil {
			return apperrors.NewInternalError("failed to create session")
		}

		result = &UserWithSession{
			User:    user,
			Session: session,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UserService) FindById(userID uint64) (*models.User, error) {
	return s.repo.FindById(userID)
}
