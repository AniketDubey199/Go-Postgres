-- name: CreateUser :one
INSERT INTO users (name, email) VALUES ($1, $2) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: GetUser :many
SELECT * FROM users;

-- name: UpdateUserByID :one
UPDATE users set
email = $1 
where user_id = $2 
RETURNING *;

-- name: DeleteUserById :one
DELETE FROM users 
where user_id = $1 RETURNING user_id;

-- name: CreateProduct :one
INSERT INTO products (product_name , unit_price , stock_qty) VALUES ($1 , $2 , $3) RETURNING *;