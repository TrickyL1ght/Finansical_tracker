package service

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/repository"
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	tokenTTL = 8 * time.Hour
)

type UserService struct {
	repository repository.Repository
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewUserService(repository repository.Repository) *UserService {
	return &UserService{repository}
}

func (s *UserService) CreateUser(user model.User) error {
	user.Password = s.GeneratePasswordHash(user.Password)
	return s.repository.CrateUser(user)
}

func (s *UserService) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}

func (s *UserService) UserAuth(login, password string) (string, error) {
	user, err := s.repository.GetUser(login, s.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
	}, user.Id})

	return token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
}

func (s *UserService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})
	if err != nil {
		return 0, err
	}
	claim, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	return claim.UserId, nil
}
