package service

import (
	"errors"
	"geobill_golang_versions/models"
	"geobill_golang_versions/repository"

	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

type Service interface {
	Register(username, password, role string) error
	Login(username, password string) (*models.User, error)
	Authenticate(username, password string) (*models.User, error)
}

type AuthService struct {
	Repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Register(username, password, role string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}

	return s.Repo.CreateUser(user)
}

func (s *AuthService) Login(username, password string) (*models.User, error) {
	return s.Authenticate(username, password)
}

func (s *AuthService) Authenticate(username, password string) (*models.User, error) {
	user, err := s.Repo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
