-- name: GetPaginatedUsers :many
SELECT * FROM users
ORDER BY user_id
LIMIT $1 OFFSET $2;

-- name: GetFilteredUsers :many
SELECT * FROM users
WHERE $1 = $2;

-- name: GetUserByID :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (passport_series, passport_number, surname, name, patronymic, address)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET passport_series = $2, passport_number = $3, surname = $4, name = $5, patronymic = $6, address = $7
WHERE user_id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;

-- name: GetUserWorklogs :many
SELECT * FROM worklogs
WHERE user_id = $1
ORDER BY start_time;

-- name: StartTask :one
INSERT INTO worklogs (user_id, title, description, start_time)
VALUES ($1, $2, $3, NOW())
RETURNING *;

-- name: StopTask :one
UPDATE worklogs
SET end_time = NOW(), 
    hours_spent = age(NOW(), start_time)
WHERE user_id = $1 AND worklog_id = $2 AND end_time IS NULL
RETURNING *;