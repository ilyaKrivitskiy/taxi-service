package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/ilyaKrivitskiy/taxi-service/models"
	"github.com/ilyaKrivitskiy/taxi-service/pkg/repository"
)

const (
	salt = "slhdf125li8w983rfh183819"
)

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

func (s *UserAuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
