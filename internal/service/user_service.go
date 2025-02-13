package service

import (
	"go-echo-experiment/internal/model"
	"go-echo-experiment/internal/repository"

	"go-echo-experiment/pkg/utils"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.Repo.GetUserByEmail(email)
}

func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
	hashPassword, salt, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	newData := &model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashPassword,
		Salt:     salt,
	}
	return s.Repo.CreateUser(newData)
}

func (s *UserService) LoginUser(user *model.User) (*model.User, error) {
	return s.Repo.LoginUser(user)
}
