// app/services/auth_service.go
package services

import (
	"errors"
	"os"
	"time"

	"product-store/app/models"
	"product-store/app/repositories"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	Repo *repositories.UserRepository
}

func NewAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Register(user *models.User) error {
	if err := user.HashPassword(user.Password); err != nil {
		return err
	}
	return s.Repo.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil || !user.CheckPassword(password) {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
