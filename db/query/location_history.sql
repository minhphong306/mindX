-- name: CreateLocationHistory :one
INSERT INTO "location_history" (user_id,
                    type,
                    location_id,
                    manual_input)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetLocationHistory :one
SELECT *
FROM "location_history"
WHERE id = $1 LIMIT 1;

-- name: ListLocationHistories :many
SELECT *
FROM "location_history"
ORDER BY id LIMIT $1
OFFSET $2;

-- name: DeleteLocationHistory :exec
DELETE
FROM "location_history"
WHERE id = $1;

