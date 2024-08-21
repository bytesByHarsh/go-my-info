-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at,
                  deleted_at, is_deleted,
                  name, phone_num, email, username,
                  profile_img, role, hashed_password)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: GetUserByUsername :one
SELECT * from users WHERE username=$1;

-- name: GetUserByEmail :one
SELECT * from users WHERE email=$1;

-- name: GetUserById :one
SELECT * from users WHERE id=$1;