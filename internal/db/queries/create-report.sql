-- name: CreateReport :one
INSERT INTO reports(user_id, team_id, report)
VALUES ($1, $2, $3)
RETURNING id;