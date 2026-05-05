package domain

import (
	"time"
)

type UserRole string

const (
	RoleUser  UserRole = "user" //client
	RoleAdmin UserRole = "admin"
)

type User struct {
	ID    int64
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	//адрес
	PasswordHash string    `json:"-"`
	Role         UserRole  `json:"role"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
