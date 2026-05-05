package domain

import (
	"context"
	"time"
)

type UserRole string

const (
	RoleUser  UserRole = "user" //client
	RoleAdmin UserRole = "admin"
)

type User struct {
	ID           string
	Name         string    `json:"username"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         UserRole  `json:"role"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByPhone(ctx context.Context, phone string) (*User, error)
}
