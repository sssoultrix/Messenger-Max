package service

import (
	"context"
	"messenger-max/user-service/internal/domain"
)

// Описание контракта для любого репозитория который будет работать с user_service
type Repository interface {
	Create(ctx context.Context, request domain.UserCreateRequest) error
	Update(ctx context.Context, request domain.UserCreateRequest) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	GetAll(ctx context.Context) ([]*domain.User, error)
	GetByLogin(ctx context.Context, login string) (*domain.User, error)
}

type UserService struct {
	repo Repository
}

func NewUserService(repo Repository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Create(ctx context.Context, request domain.UserCreateRequest) error {
	return u.repo.Create(ctx, request)
}

func (u *UserService) Update(ctx context.Context, request domain.UserCreateRequest) error {
	return u.repo.Update(ctx, request)
}

func (u *UserService) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *UserService) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *UserService) GetAll(ctx context.Context) ([]*domain.User, error) {
	return u.repo.GetAll(ctx)
}

func (u *UserService) GetByLogin(ctx context.Context, login string) (*domain.User, error) {
	return u.repo.GetByLogin(ctx, login)
}
