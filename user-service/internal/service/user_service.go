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
