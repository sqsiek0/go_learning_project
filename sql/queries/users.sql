-- name: CreateUser :one

INSERT INTO
    users (
        id, createdat, updatedat, name, surname, api_key
    )
VALUES (
        $1, $2, $3, $4, $5, encode(
            sha256(random()::text::bytea), 'hex'
        )
    )
RETURNING
    *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1;