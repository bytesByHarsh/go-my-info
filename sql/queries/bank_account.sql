-- name: CreateBankAccount :one
INSERT INTO bank_accounts(
    id, created_at, updated_at, deleted_at, is_deleted, is_active,
    bank_id, user_id, account_number, account_type,
    name, balance, currency)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING *;


-- name: GetBankAccountById :one
SELECT * from bank_accounts WHERE id=$1 AND is_deleted = false;

-- name: GetAllBankAccount :many
SELECT *
FROM
    bank_accounts
WHERE is_deleted = false
ORDER BY
    created_at ASC
LIMIT $1 OFFSET $2;

-- name: GetBankAccountCount :one
SELECT COUNT(*) FROM bank_accounts WHERE is_deleted=false;

-- name: GetUserBankAccounts :many
SELECT *
FROM
    bank_accounts
WHERE is_deleted = false AND user_id=$3
ORDER BY
    created_at ASC
LIMIT $1 OFFSET $2;

-- name: GetUserBankAccountCount :one
SELECT COUNT(*) FROM bank_accounts
WHERE is_deleted=false AND user_id=$1;


-- name: UpdateBankAccount :exec
UPDATE bank_accounts
SET updated_at = $2,
    account_number = $3,
    account_type = $4,
    is_active = $5,
    name = $6,
    balance = $7,
    currency = $8
WHERE id = $1 AND is_deleted=false
RETURNING *;


-- name: DeleteBankAccount :exec
UPDATE bank_accounts
SET deleted_at = $2,
    is_deleted = true,
    updated_at = $3
WHERE id = $1
RETURNING *;

-- name: HardDeleteBankAccount :exec
DELETE FROM bank_accounts
WHERE id = $1;