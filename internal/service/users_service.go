package service

import (
	"context"
	"errors"
	"koda-b8-backend1/internal/domain"
	"koda-b8-backend1/internal/repository"
)

type UserService struct { 
  repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService { 
  return &UserService{repository: repository}
}

func (s *UserService) Create(data *domain.CreateUserRequest, ctx context.Context) (*domain.User, error) {
    if len(data.Password) < 8 { 
      return &domain.User{}, errors.New("Password Harus miniimal 8 Karakter")
    }
    return s.repository.Create(data, ctx)
}

func (s *UserService) GetUsers(ctx context.Context) (*[]domain.User, error) { 
  return s.repository.FindAll(ctx)
}

func (s *UserService) Login(data *domain.LoginRequest, ctx context.Context) (*domain.User, error) {
	user, err := s.repository.FindByEmail(data.Email, ctx)
	if err != nil {
		return nil, err
	}
	if user.Password != data.Password {
		return nil, errors.New("password salah")
	}
	return user, nil
}