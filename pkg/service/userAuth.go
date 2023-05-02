package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ilyaKrivitskiy/taxi-service/models"
	"github.com/ilyaKrivitskiy/taxi-service/pkg/repository"
	"time"
)

const (
	salt       = "slhdf125li8w983rfh183819"
	tokenTTL   = 12 * time.Hour
	signingKey = "sjefk1k2kqug72ig378rq"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId uint `json:"user_id"`
}

type UserAuthService struct {
	repo repository.UserAuthorization
}

func NewUserAuthService(repo repository.UserAuthorization) *UserAuthService {
	return &UserAuthService{repo: repo}
}

func (s *UserAuthService) CreateUser(user models.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *UserAuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *UserAuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
