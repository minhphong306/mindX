-- name: CreateUser :one
INSERT INTO "user" (name,
                    permanent_address,
                    current_address,
                    current_status)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUser :one
SELECT *
FROM "user"
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM "user"
ORDER BY id LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE "user"
SET permanent_address = $2,
    current_address   = $3,
    current_status    = $4,
    name              = $5

WHERE id = $1 RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM "user"
WHERE id = $1;

