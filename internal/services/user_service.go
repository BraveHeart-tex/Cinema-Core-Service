package services

import (
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo           *repositories.UserRepository
	sessionService *SessionService
}

func NewUserService(repo *repositories.UserRepository, sessionService *SessionService) *UserService {
	return &UserService{
		repo:           repo,
		sessionService: sessionService,
	}
}

type CreateUserData struct {
	Name     string
	Surname  string
	Email    string
	Password string
}

type UserWithSession struct {
	User    *models.User
	Session *models.SessionWithToken
}

func (s *UserService) CreateUser(data CreateUserData) (*UserWithSession, error) {
	existing, _ := s.repo.FindByEmail(data.Email)
	if existing != nil {
		return nil, NewConflict("user already exists with the given email")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, NewInternalError("failed to hash password")
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
		return nil, NewInternalError("failed to create user")
	}

	// Create session for the new user
	session, err := s.sessionService.CreateSession(user.Id)
	if err != nil {
		return nil, NewInternalError("failed to create session")
	}

	return &UserWithSession{
		User:    createdUser,
		Session: session,
	}, nil
}

type SignInData struct {
	Email    string
	Password string
}

func (s *UserService) SignIn(data SignInData) (*UserWithSession, error) {
	user, err := s.repo.FindByEmail(data.Email)
	if err != nil {
		return nil, NewInternalError("failed to fetch user")
	}

	// TODO: Add better nil / not found checks here
	if user == nil {
		return nil, NewUnauthorized("invalid email or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(data.Password)) != nil {
		return nil, NewUnauthorized("invalid email or password")
	}

	session, err := s.sessionService.CreateSession(user.Id)
	if err != nil {
		return nil, NewInternalError("failed to create session")
	}

	return &UserWithSession{
		User:    user,
		Session: session,
	}, nil
}
