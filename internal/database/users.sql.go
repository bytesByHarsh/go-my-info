// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at,
                  deleted_at, is_deleted,
                  name, phone_num, email,
                  profile_img, role, hashed_password)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id, created_at, updated_at, deleted_at, is_deleted, name, phone_num, email, profile_img, role, hashed_password
`

type CreateUserParams struct {
	ID             uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime
	IsDeleted      bool
	Name           string
	PhoneNum       sql.NullString
	Email          string
	ProfileImg     sql.NullString
	Role           int32
	HashedPassword sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.IsDeleted,
		arg.Name,
		arg.PhoneNum,
		arg.Email,
		arg.ProfileImg,
		arg.Role,
		arg.HashedPassword,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.IsDeleted,
		&i.Name,
		&i.PhoneNum,
		&i.Email,
		&i.ProfileImg,
		&i.Role,
		&i.HashedPassword,
	)
	return i, err
}
