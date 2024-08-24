-- name: CreateCard :one
INSERT INTO cards(
    id, created_at, updated_at, deleted_at, is_deleted,is_active,
    bank_id, user_id, bank_account_id, name, nickname,
    number, type, expiration_date, cvv, total_limit, bill_date)
VALUES ($1, $2, $3, $4, $5, $6, $7,
        $8, $9, $10, $11, $12, $13,
        $14, $15, $16, $17)
RETURNING *;


-- name: GetCardById :one
SELECT * from cards WHERE id=$1 AND is_deleted = false;

-- name: GetAllCard :many
SELECT *
FROM
    cards
WHERE is_deleted = false
ORDER BY
    name ASC
LIMIT $1 OFFSET $2;

-- name: GetCardCount :one
SELECT COUNT(*) FROM cards WHERE is_deleted=false;

-- name: GetUserCards :many
SELECT *
FROM
    cards
WHERE is_deleted = false AND user_id=$3
ORDER BY
    name ASC
LIMIT $1 OFFSET $2;

-- name: GetUserCardCount :one
SELECT COUNT(*) FROM cards WHERE is_deleted=false AND user_id=$1;


-- name: UpdateCard :exec
UPDATE cards
SET updated_at = $2,
    name = $3,
    nickname = $4,
    is_active = $5,
    number = $6,
    type = $7,
    expiration_date = $8,
    cvv = $9,
    name = $10,
    nickname = $11
WHERE id = $1 AND is_deleted=false
RETURNING *;


-- name: DeleteCard :exec
UPDATE cards
SET deleted_at = $2,
    is_deleted = true,
    updated_at = $3
WHERE id = $1
RETURNING *;

-- name: HardDeleteCard :exec
DELETE FROM cards
WHERE id = $1;