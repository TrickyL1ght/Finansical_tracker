package user

import (
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	salt       = os.Getenv("SALT")
	tokenTTL   = 12 * time.Hour
	signingKey = os.Getenv("SIGNING_KEY")
)

type UserService interface {
	CreateUser(user User) error
	UserAuth(login, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type service struct {
	repository UserRepository
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewUserService(repository UserRepository) UserService {
	return &service{repository}
}

func (s *service) CreateUser(user User) error {
	user.Password = s.GeneratePasswordHash(user.Password)
	return s.repository.CrateUser(user)
}

func (s *service) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *service) UserAuth(login, password string) (string, error) {
	user, err := s.repository.GetUser(login, s.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
	}, user.Id})

	return token.SignedString([]byte(signingKey))
}

func (s *service) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
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
