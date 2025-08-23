package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"messenger-max/user-service/internal/domain"
	"messenger-max/user-service/pkg/hash"
	"messenger-max/user-service/pkg/logger"
)

type UserPostgres struct {
	pool *pgxpool.Pool
}

func NewUserPostgres(pool *pgxpool.Pool) *UserPostgres {
	return &UserPostgres{pool: pool}
}

func (u UserPostgres) Create(ctx context.Context, request domain.UserCreateRequest) error {
	query := `INSERT INTO users (login, password_hash) VALUES ($1, $2)`

	hashedPassword, err := hash.HashPassword(request.Password)
	if err != nil {
		logger.Log.Error("failed to hash password", "error", err)
		return err
	}
	_, err = u.pool.Exec(ctx, query, request.Login, hashedPassword)
	if err != nil {
		logger.Log.Error("failed to insert user", "error", err)
		return err
	}
	logger.Log.Info("User created", "Login", request.Login)
	return nil
}

func (u *UserPostgres) Update(ctx context.Context, request domain.UserCreateRequest) error {
	query := `UPDATE users SET login = $1, password_hash = $2 WHERE id = $3`
	if Password := request.Password; Password != "" {
		hashedPassword, err := hash.HashPassword(Password)
		if err != nil {
			logger.Log.Error("failed to hash password", "error", err)
			return err
		}
		_, err = u.pool.Exec(ctx, query, request.Login, hashedPassword, request.ID)
		if err != nil {
			logger.Log.Error("failed to update user", "error", err)
			return err
		}
	}
	_, err := u.pool.Exec(ctx, query, request.Login, request.ID)
	if err != nil {
		logger.Log.Error("failed to insert user", "error", err)
		return err
	}
	logger.Log.Info("User data updated", "Login", request.Login)
	return nil
}

func (u *UserPostgres) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := u.pool.Exec(ctx, query, id)

	if err != nil {
		logger.Log.Error("failed to delete user", "error", err)
		return err
	}
	logger.Log.Info("User deleted", "id", id)
	return nil
}

func (u *UserPostgres) GetByID(ctx context.Context, id int64) (*domain.UserResponse, error) {
	query := `SELECT login, name FROM users WHERE id = $1`
	row := u.pool.QueryRow(ctx, query, id)
	var user domain.UserResponse
	if err := row.Scan(&user.Login, &user.Name); err != nil {
		logger.Log.Error("failed to fetch user", "error", err)
		return nil, err
	}
	logger.Log.Info("User data retrieved", "Login", user.Login)
	return &user, nil
}

func (u *UserPostgres) GetAll(ctx context.Context) ([]domain.UserResponse, error) {
	query := `SELECT id, login, name FROM users`
	rows, err := u.pool.Query(ctx, query)
	if err != nil {
		logger.Log.Error("failed to fetch users", "error", err)
		return nil, err
	}
	defer rows.Close()
	var users []domain.UserResponse
	for rows.Next() {
		var user domain.UserResponse

		if err := rows.Scan(&user.ID, &user.Login, &user.Name); err != nil {
			logger.Log.Error("failed to fetch users", "error", err)
			return nil, err
		}
		users = append(users, user)
		logger.Log.Info("User data retrieved", "Login", user.Login)
	}
	logger.Log.Info("Users data retrieved", "Count", len(users))
	return users, nil
}

func (u *UserPostgres) GetByLogin(ctx context.Context, login string) (domain.UserResponse, error) {
	query := `SELECT id, login, name FROM users WHERE login = $1`
	row := u.pool.QueryRow(ctx, query, login)
	var user domain.UserResponse
	if err := row.Scan(&user.ID, &user.Login, &user.Name); err != nil {
		logger.Log.Error("failed to fetch user", "error", err)
		return domain.UserResponse{}, err
	}
	logger.Log.Info("User data retrieved", "Login", user.Login)
	return user, nil
}
