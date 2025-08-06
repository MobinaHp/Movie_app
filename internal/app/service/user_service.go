package service

import (
	"WEB1/internal/app/repository"
	"WEB1/internal/domain"
)

type userService struct{
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) RegisterUser(name, email, password string) (domain.User, error) {
	user := domain.User{
		Name:     name,
		Email:    email,
		HashedPassword: password,
	}
	return s.repo.Create(user)
}

func (s *userService) GetUserByID(id uint) (domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) ListUsers() ([]domain.User, error) {
	return s.repo.List()
}