package service

import (
	"context"
	"github.com/thisgleammm/mantis-backend/internal/domain"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ListUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.List(ctx)
}

func (s *UserService) FindUserByID(ctx context.Context, id string) (domain.User, error) {
	return s.repo.FindByID(ctx, id)
}
