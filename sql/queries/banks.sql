-- name: CreateBank :one
INSERT INTO banks(
    id, created_at, updated_at, deleted_at, is_deleted,
    name, contact_phone, contact_email, address,
    type, established_year)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;


-- name: GetBankById :one
SELECT * from banks WHERE id=$1 AND is_deleted = false;

-- name: GetAllBank :many
SELECT *
FROM
    banks
WHERE is_deleted = false
ORDER BY
    name ASC
LIMIT $1 OFFSET $2;

-- name: GetBankCount :one
SELECT COUNT(*) FROM banks WHERE is_deleted=false;

-- name: UpdateBank :exec
UPDATE banks
SET updated_at = $2,
    name = $3,
    contact_phone = $4,
    contact_email = $5,
    address = $6,
    type = $7
WHERE id = $1 AND is_deleted=false
RETURNING *;


-- name: DeleteBank :exec
UPDATE banks
SET deleted_at = $2,
    is_deleted = true,
    updated_at = $3
WHERE id = $1
RETURNING *;

-- name: HardDeleteBank :exec
DELETE FROM banks
WHERE id = $1;