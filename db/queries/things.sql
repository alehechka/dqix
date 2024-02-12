-- name: GetThingById :one
SELECT * FROM things
WHERE id = ? LIMIT 1;