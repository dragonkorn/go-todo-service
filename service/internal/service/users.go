package service

import (
	"fmt"
	"service/internal/config"
	"service/internal/model"
	"service/internal/repository"
)

type UserService struct {
	Config  *config.Configuration
	UserRepo *repository.UserRepository
}

func NewUserService(
	cf *config.Configuration,
	userRp *repository.UserRepository,
) *UserService {
	return &UserService{
		Config: cf,
		UserRepo: userRp,
	}
}

func(s *UserService) CreateUser(user *model.Users) error{
	if err := s.UserRepo.CreateOrUpdateUser(user); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}