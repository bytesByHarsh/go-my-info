-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at,
                  deleted_at, is_deleted,
                  name, phone_num, email, username,
                  profile_img, role, hashed_password, is_active)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, TRUE)
RETURNING *;

-- name: GetUserByUsername :one
SELECT * from users WHERE username=$1 AND is_deleted = false;

-- name: GetUserByEmail :one
SELECT * from users WHERE email=$1 AND is_deleted = false;

-- name: GetUserById :one
SELECT * from users WHERE id=$1 AND is_deleted = false;

-- name: GetAllUsers :many
SELECT *
FROM
    users
WHERE is_deleted = false
ORDER BY
    name ASC
LIMIT $1 OFFSET $2;

-- name: GetUserCount :one
SELECT COUNT(*) FROM users WHERE is_deleted=false;

-- name: UpdateUser :exec
UPDATE users
SET updated_at = $2,
    name = $3,
    phone_num = $4,
    email = $5,
    username = $6,
    profile_img = $7,
    role = $8
WHERE id = $1 AND is_deleted=false
RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users
SET hashed_password=$2,
    updated_at = $3
WHERE id = $1 AND is_deleted = false
RETURNING *;

-- name: DeleteUser :exec
UPDATE users
SET deleted_at = $2,
    is_deleted = true,
    updated_at = $3
WHERE id = $1
RETURNING *;

-- name: HardDeleteUser :exec
DELETE FROM users
WHERE id = $1;