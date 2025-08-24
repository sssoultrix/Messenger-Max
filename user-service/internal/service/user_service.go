package service

import (
	"context"
	"messenger-max/user-service/internal/domain"
	"messenger-max/user-service/pkg/hash"
	"messenger-max/user-service/pkg/logger"
)

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
	var err error
	request.Password, err = hash.HashPassword(request.Password)
	if err != nil {
		logger.Log.Error("failed to hash password", "error", err)
		return err
	}
	return u.repo.Create(ctx, request)
}

func (u *UserService) Update(ctx context.Context, request domain.UserCreateRequest) error {
	if Password := request.Password; Password != "" {
		hashedPassword, err := hash.HashPassword(Password)
		if err != nil {
			logger.Log.Error("failed to hash password", "error", err)
			return err
		}
		request.Password = hashedPassword
	}
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
