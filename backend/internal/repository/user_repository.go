package repository

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID
	Email string
	PasswordHash string
	Name string
	Status string
}

type UserRepository interface {
	Create(ctx context.Context, user User) error

	FindByEmail(
		ctx context.Context,
		email string,
	) (*User, error)

	FindByID(
		ctx context.Context,
		id uuid.UUID,
	) (*User, error)
}