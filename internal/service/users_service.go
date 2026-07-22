package service

import (
	"context"
	"errors"
	"fmt"
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

func (s *UserService) GetUserByID(id int64, ctx context.Context) (*domain.User, error) { 
  return s.repository.FindByID(id, ctx)
}

func (s *UserService) DeleteUser(id int64, ctx context.Context) (error) { 
  user, err := s.repository.FindByID(id, ctx) 
  fmt.Print(user)

  if err != nil { 
    return nil
  }
  return s.repository.Delete(int64(user.Id), ctx)
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

func (s *UserService) Patch(id int64, req *domain.PatchUserRequest, ctx context.Context) (*domain.User, error) { 
  _, err := s.repository.FindByID(id, ctx)
  if err != nil {
    return nil, err
  }
  return s.repository.Patch(id, req, ctx)
}

func (s *UserService) UploadPictureProfile(id int64, req *domain.UploadPicturesProfileRequest, ctx context.Context) (*domain.User, error) {
  _, err := s.repository.FindByID(id, ctx)
  if err != nil {
    return nil, err
  }
  return s.repository.UploadPictureProfile(id, req, ctx)
}