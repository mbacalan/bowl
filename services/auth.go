package services

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/mbacalan/bowl/models"
	"github.com/mbacalan/bowl/services/internal"
)

type AuthService struct {
	Logger     *slog.Logger
	Repository models.UserRepository
	Hash       *internal.Argon2idHash
}

func NewAuthService(logger *slog.Logger, repo models.UserRepository, hash *internal.Argon2idHash) *AuthService {
	return &AuthService{
		Logger:     logger,
		Repository: repo,
		Hash:       internal.NewArgon2idHash(1, 32, 64*1024, 32, 256),
	}
}

func (s *AuthService) Signup(name string, password string) (models.User, error) {
	userPassword := []byte(password)
	hashSalt, err := s.Hash.GenerateHash(userPassword, []byte(os.Getenv("PW_SALT")))

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return models.User{}, err
	}

	user, err := s.Repository.Create(name, hashSalt.Hash)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return models.User{}, err
	}

	return user, nil
}

func (s *AuthService) Login(name string, attempt string) (models.User, error) {
	user, err := s.Repository.Get(name)

	if err != nil {
		s.Logger.Error("Error logging in", err)
		return models.User{}, err
	}

	err = s.Hash.CompareHashAndPassword(user.Password, []byte(os.Getenv("PW_SALT")), []byte(attempt))

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return models.User{}, err
	}

	fmt.Println("argon2IDHash Password and Hash match")

	return user, nil
}
