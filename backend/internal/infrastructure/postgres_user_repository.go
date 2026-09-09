package infrastructure

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dirgefigure-source/enterprise-flow/internal/repository"
)

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(
	db *pgxpool.Pool,
) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) Create(
	ctx context.Context,
	user repository.User,
) error {
	_, err := r.db.Exec(
		ctx,
		`
		INSERT INTO users (
			id,
			email,
			password_hash,
			name,
			status
		)
		VALUES ($1, $2, $3, $4, $5)
		`,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.Name,
		user.Status,
	)

	return err
}

func (r *PostgresUserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*repository.User, error) {
	var user repository.User

	err := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			email,
			password_hash,
			name,
			status
		FROM users
		WHERE email = $1
		`,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.Status,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*repository.User, error) {
	var user repository.User

	err := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			email,
			password_hash,
			name,
			status
		FROM users
		WHERE id = $1
		`,
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.Status,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}