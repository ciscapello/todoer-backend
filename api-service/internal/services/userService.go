package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/ciscapello/api-service/internal/domain/model"
	"github.com/ciscapello/api-service/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidPassword   = errors.New("invalid password")
	ErrUserDoesNotExists = errors.New("user does not exists")
)

type UserService struct {
	usersRepo    *repository.UsersRepo
	sessionsRepo *repository.SessionsRepo
}

func NewUserService(usersRepo *repository.UsersRepo, sessionsRepo *repository.SessionsRepo) *UserService {
	return &UserService{usersRepo: usersRepo, sessionsRepo: sessionsRepo}
}

func (s *UserService) CreateUser(email string, password string) (int, string) {
	isUserExists := s.usersRepo.CheckUserIsExists(email)
	if isUserExists {
		return 0, "this username is taken"
	}

	id, errorMessage := s.usersRepo.CreateUser(email, password)
	return id, errorMessage
}

func (s *UserService) CreateSession(userId int, token string, expired_at time.Time) int {
	id := s.usersRepo.CreateSession(userId, token, expired_at)
	return id
}

func (s *UserService) SignInUser(email string, password string) (model.User, error) {

	user, err := s.usersRepo.GetByEmail(email)
	if err != nil || user.Email == "" {
		return model.User{}, ErrUserDoesNotExists
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	fmt.Println(err)
	if err != nil {
		return model.User{}, ErrInvalidPassword
	}

	usr, err := s.usersRepo.GetByEmail(email)
	fmt.Printf("%v", user)
	if err != nil {
		fmt.Println(err)
	}

	return usr, nil
}
