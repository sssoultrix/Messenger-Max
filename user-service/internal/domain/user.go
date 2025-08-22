package domain

import "time"

type User struct {
	ID           int64     `json:"id"`
	Login        string    `json:"Login"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserCreateRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
