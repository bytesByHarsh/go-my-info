// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: banks.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createBank = `-- name: CreateBank :one
INSERT INTO banks(
    id, created_at, updated_at, deleted_at, is_deleted,
    name, contact_phone, contact_email, address,
    type, established_year)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id, created_at, updated_at, deleted_at, is_deleted, name, contact_phone, contact_email, address, type, established_year
`

type CreateBankParams struct {
	ID              uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime
	IsDeleted       bool
	Name            string
	ContactPhone    sql.NullString
	ContactEmail    string
	Address         sql.NullString
	Type            BankType
	EstablishedYear int32
}

func (q *Queries) CreateBank(ctx context.Context, arg CreateBankParams) (Bank, error) {
	row := q.db.QueryRowContext(ctx, createBank,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.IsDeleted,
		arg.Name,
		arg.ContactPhone,
		arg.ContactEmail,
		arg.Address,
		arg.Type,
		arg.EstablishedYear,
	)
	var i Bank
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.IsDeleted,
		&i.Name,
		&i.ContactPhone,
		&i.ContactEmail,
		&i.Address,
		&i.Type,
		&i.EstablishedYear,
	)
	return i, err
}

const deleteBank = `-- name: DeleteBank :exec
UPDATE banks
SET deleted_at = $2,
    is_deleted = true,
    updated_at = $3
WHERE id = $1
RETURNING id, created_at, updated_at, deleted_at, is_deleted, name, contact_phone, contact_email, address, type, established_year
`

type DeleteBankParams struct {
	ID        uuid.UUID
	DeletedAt sql.NullTime
	UpdatedAt time.Time
}

func (q *Queries) DeleteBank(ctx context.Context, arg DeleteBankParams) error {
	_, err := q.db.ExecContext(ctx, deleteBank, arg.ID, arg.DeletedAt, arg.UpdatedAt)
	return err
}

const getAllBank = `-- name: GetAllBank :many
SELECT id, created_at, updated_at, deleted_at, is_deleted, name, contact_phone, contact_email, address, type, established_year
FROM
    banks
WHERE is_deleted = false
ORDER BY
    name ASC
LIMIT $1 OFFSET $2
`

type GetAllBankParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetAllBank(ctx context.Context, arg GetAllBankParams) ([]Bank, error) {
	rows, err := q.db.QueryContext(ctx, getAllBank, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Bank
	for rows.Next() {
		var i Bank
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.IsDeleted,
			&i.Name,
			&i.ContactPhone,
			&i.ContactEmail,
			&i.Address,
			&i.Type,
			&i.EstablishedYear,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBankById = `-- name: GetBankById :one
SELECT id, created_at, updated_at, deleted_at, is_deleted, name, contact_phone, contact_email, address, type, established_year from banks WHERE id=$1 AND is_deleted = false
`

func (q *Queries) GetBankById(ctx context.Context, id uuid.UUID) (Bank, error) {
	row := q.db.QueryRowContext(ctx, getBankById, id)
	var i Bank
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.IsDeleted,
		&i.Name,
		&i.ContactPhone,
		&i.ContactEmail,
		&i.Address,
		&i.Type,
		&i.EstablishedYear,
	)
	return i, err
}

const getBankByName = `-- name: GetBankByName :one
SELECT id, created_at, updated_at, deleted_at, is_deleted, name, contact_phone, contact_email, address, type, established_year from banks WHERE name=$1 AND is_deleted = false LIMIT 1
`

func (q *Queries) GetBankByName(ctx context.Context, name string) (Bank, error) {
	row := q.db.QueryRowContext(ctx, getBankByName, name)
	var i Bank
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.IsDeleted,
		&i.Name,
		&i.ContactPhone,
		&i.ContactEmail,
		&i.Address,
		&i.Type,
		&i.EstablishedYear,
	)
	return i, err
}

const getBankCount = `-- name: GetBankCount :one
SELECT COUNT(*) FROM banks WHERE is_deleted=false
`

func (q *Queries) GetBankCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getBankCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const hardDeleteBank = `-- name: HardDeleteBank :exec
DELETE FROM banks
WHERE id = $1
`

func (q *Queries) HardDeleteBank(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, hardDeleteBank, id)
	return err
}

const updateBank = `-- name: UpdateBank :exec
UPDATE banks
SET updated_at = $2,
    name = $3,
    contact_phone = $4,
    contact_email = $5,
    address = $6,
    type = $7
WHERE id = $1 AND is_deleted=false
RETURNING id, created_at, updated_at, deleted_at, is_deleted, name, contact_phone, contact_email, address, type, established_year
`

type UpdateBankParams struct {
	ID           uuid.UUID
	UpdatedAt    time.Time
	Name         string
	ContactPhone sql.NullString
	ContactEmail string
	Address      sql.NullString
	Type         BankType
}

func (q *Queries) UpdateBank(ctx context.Context, arg UpdateBankParams) error {
	_, err := q.db.ExecContext(ctx, updateBank,
		arg.ID,
		arg.UpdatedAt,
		arg.Name,
		arg.ContactPhone,
		arg.ContactEmail,
		arg.Address,
		arg.Type,
	)
	return err
}
