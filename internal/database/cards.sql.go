// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: cards.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createCard = `-- name: CreateCard :one
INSERT INTO cards(
    id, created_at, updated_at, deleted_at, is_deleted,is_active,
    bank_id, user_id, bank_account_id, name, nickname,
    number, type, exp_month, exp_year, cvv, total_limit, bill_date)
VALUES ($1, $2, $3, $4, $5, $6, $7,
        $8, $9, $10, $11, $12, $13,
        $14, $15, $16, $17, $18)
RETURNING id, created_at, updated_at, deleted_at, is_deleted, is_active, user_id, bank_id, bank_account_id, name, nickname, number, type, exp_month, exp_year, cvv, total_limit, bill_date
`

type CreateCardParams struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime
	IsDeleted     bool
	IsActive      bool
	BankID        uuid.UUID
	UserID        uuid.UUID
	BankAccountID uuid.NullUUID
	Name          string
	Nickname      string
	Number        string
	Type          CardType
	ExpMonth      int32
	ExpYear       int32
	Cvv           string
	TotalLimit    string
	BillDate      int32
}

func (q *Queries) CreateCard(ctx context.Context, arg CreateCardParams) (Card, error) {
	row := q.db.QueryRowContext(ctx, createCard,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.IsDeleted,
		arg.IsActive,
		arg.BankID,
		arg.UserID,
		arg.BankAccountID,
		arg.Name,
		arg.Nickname,
		arg.Number,
		arg.Type,
		arg.ExpMonth,
		arg.ExpYear,
		arg.Cvv,
		arg.TotalLimit,
		arg.BillDate,
	)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.IsDeleted,
		&i.IsActive,
		&i.UserID,
		&i.BankID,
		&i.BankAccountID,
		&i.Name,
		&i.Nickname,
		&i.Number,
		&i.Type,
		&i.ExpMonth,
		&i.ExpYear,
		&i.Cvv,
		&i.TotalLimit,
		&i.BillDate,
	)
	return i, err
}

const deleteCard = `-- name: DeleteCard :exec
UPDATE cards
SET deleted_at = $2,
    is_deleted = true,
    updated_at = $3
WHERE id = $1
RETURNING id, created_at, updated_at, deleted_at, is_deleted, is_active, user_id, bank_id, bank_account_id, name, nickname, number, type, exp_month, exp_year, cvv, total_limit, bill_date
`

type DeleteCardParams struct {
	ID        uuid.UUID
	DeletedAt sql.NullTime
	UpdatedAt time.Time
}

func (q *Queries) DeleteCard(ctx context.Context, arg DeleteCardParams) error {
	_, err := q.db.ExecContext(ctx, deleteCard, arg.ID, arg.DeletedAt, arg.UpdatedAt)
	return err
}

const getAllCard = `-- name: GetAllCard :many
SELECT id, created_at, updated_at, deleted_at, is_deleted, is_active, user_id, bank_id, bank_account_id, name, nickname, number, type, exp_month, exp_year, cvv, total_limit, bill_date
FROM
    cards
WHERE is_deleted = false
ORDER BY
    name ASC
LIMIT $1 OFFSET $2
`

type GetAllCardParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetAllCard(ctx context.Context, arg GetAllCardParams) ([]Card, error) {
	rows, err := q.db.QueryContext(ctx, getAllCard, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Card
	for rows.Next() {
		var i Card
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.IsDeleted,
			&i.IsActive,
			&i.UserID,
			&i.BankID,
			&i.BankAccountID,
			&i.Name,
			&i.Nickname,
			&i.Number,
			&i.Type,
			&i.ExpMonth,
			&i.ExpYear,
			&i.Cvv,
			&i.TotalLimit,
			&i.BillDate,
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

const getCardById = `-- name: GetCardById :one
SELECT id, created_at, updated_at, deleted_at, is_deleted, is_active, user_id, bank_id, bank_account_id, name, nickname, number, type, exp_month, exp_year, cvv, total_limit, bill_date from cards WHERE id=$1 AND is_deleted = false
`

func (q *Queries) GetCardById(ctx context.Context, id uuid.UUID) (Card, error) {
	row := q.db.QueryRowContext(ctx, getCardById, id)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.IsDeleted,
		&i.IsActive,
		&i.UserID,
		&i.BankID,
		&i.BankAccountID,
		&i.Name,
		&i.Nickname,
		&i.Number,
		&i.Type,
		&i.ExpMonth,
		&i.ExpYear,
		&i.Cvv,
		&i.TotalLimit,
		&i.BillDate,
	)
	return i, err
}

const getCardCount = `-- name: GetCardCount :one
SELECT COUNT(*) FROM cards WHERE is_deleted=false
`

func (q *Queries) GetCardCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getCardCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getUserCardCount = `-- name: GetUserCardCount :one
SELECT COUNT(*) FROM cards WHERE is_deleted=false AND user_id=$1
`

func (q *Queries) GetUserCardCount(ctx context.Context, userID uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, getUserCardCount, userID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getUserCards = `-- name: GetUserCards :many
SELECT id, created_at, updated_at, deleted_at, is_deleted, is_active, user_id, bank_id, bank_account_id, name, nickname, number, type, exp_month, exp_year, cvv, total_limit, bill_date
FROM
    cards
WHERE is_deleted = false AND user_id=$3
ORDER BY
    name ASC
LIMIT $1 OFFSET $2
`

type GetUserCardsParams struct {
	Limit  int32
	Offset int32
	UserID uuid.UUID
}

func (q *Queries) GetUserCards(ctx context.Context, arg GetUserCardsParams) ([]Card, error) {
	rows, err := q.db.QueryContext(ctx, getUserCards, arg.Limit, arg.Offset, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Card
	for rows.Next() {
		var i Card
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.IsDeleted,
			&i.IsActive,
			&i.UserID,
			&i.BankID,
			&i.BankAccountID,
			&i.Name,
			&i.Nickname,
			&i.Number,
			&i.Type,
			&i.ExpMonth,
			&i.ExpYear,
			&i.Cvv,
			&i.TotalLimit,
			&i.BillDate,
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

const hardDeleteCard = `-- name: HardDeleteCard :exec
DELETE FROM cards
WHERE id = $1
`

func (q *Queries) HardDeleteCard(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, hardDeleteCard, id)
	return err
}

const updateCard = `-- name: UpdateCard :exec
UPDATE cards
SET updated_at = $2,
    name = $3,
    nickname = $4,
    is_active = $5,
    number = $6,
    type = $7,
    exp_month = $8,
    exp_year = $9,
    cvv = $10
WHERE id = $1 AND is_deleted=false
RETURNING id, created_at, updated_at, deleted_at, is_deleted, is_active, user_id, bank_id, bank_account_id, name, nickname, number, type, exp_month, exp_year, cvv, total_limit, bill_date
`

type UpdateCardParams struct {
	ID        uuid.UUID
	UpdatedAt time.Time
	Name      string
	Nickname  string
	IsActive  bool
	Number    string
	Type      CardType
	ExpMonth  int32
	ExpYear   int32
	Cvv       string
}

func (q *Queries) UpdateCard(ctx context.Context, arg UpdateCardParams) error {
	_, err := q.db.ExecContext(ctx, updateCard,
		arg.ID,
		arg.UpdatedAt,
		arg.Name,
		arg.Nickname,
		arg.IsActive,
		arg.Number,
		arg.Type,
		arg.ExpMonth,
		arg.ExpYear,
		arg.Cvv,
	)
	return err
}
