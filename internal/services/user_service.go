package services

import (
	"errors"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

type CreateUserData struct {
	Name     string
	Surname  string
	Email    string
	Password string
}

func (s *UserService) CreateUser(data CreateUserData) (*models.User, error) {
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

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
