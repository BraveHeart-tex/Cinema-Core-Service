package services

import (
	"errors"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo          *repositories.UserRepository
	sessionService *SessionService
}

func NewUserService(repo *repositories.UserRepository, sessionService *SessionService) *UserService {
	return &UserService{
		repo:          repo,
		sessionService: sessionService,
	}
}

type CreateUserData struct {
	Name     string
	Surname  string
	Email    string
	Password string
}

type CreateUserResult struct {
	User    *models.User
	Session *models.SessionWithToken
}

func (s *UserService) CreateUser(data CreateUserData) (*CreateUserResult, error) {
	existing, _ := s.repo.GetByEmail(data.Email)
	if existing != nil {
		return nil, errors.New("user already exists with the given email")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:           data.Name,
		Surname:        data.Surname,
		Email:          data.Email,
		HashedPassword: string(hashed),
		Role:           models.UserRole,
	}

	createdUser, err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	// Create session for the new user
	session, err := s.sessionService.CreateSession()
	if err != nil {
		return nil, err
	}

	return &CreateUserResult{
		User:    createdUser,
		Session: session,
	}, nil
}
