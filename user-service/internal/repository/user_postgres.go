package repository

import (
	"context"
	"messenger-max/user-service/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserPostgres struct {
	pool *pgxpool.Pool
}

func NewUserPostgres(pool *pgxpool.Pool) *UserPostgres {
	return &UserPostgres{pool: pool}
}

func (u UserPostgres) Create(ctx context.Context, request domain.UserCreateRequest) error {
	panic("implement me")
}

func (u UserPostgres) Update(ctx context.Context, request domain.UserCreateRequest) error {
	//TODO implement me
	panic("implement me")
}

func (u UserPostgres) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (u UserPostgres) GetByID(ctx context.Context, id int64) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserPostgres) GetAll(ctx context.Context) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserPostgres) GetByLogin(ctx context.Context, login string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}
