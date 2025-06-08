package usecase

import "github.com/santigorbe/hexagonal_arq/internal/domain"

type UserServiceImpl struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) GetAllUsers() []domain.User {
	return s.repo.FindAll()
}

func (s *UserServiceImpl) CreateUser(u domain.User) domain.User {
	return s.repo.Save(u)
}
