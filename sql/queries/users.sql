-- name: CreateUser :one

INSERT INTO
    users (
        id, createdat, updatedat, name, surname
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING
    *;