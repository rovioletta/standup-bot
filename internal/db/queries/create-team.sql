-- name: CreateTeam :one
INSERT INTO
  teams(
    team_name,
    manager_id,
    team_members,
    channel_id,
    notification_hour
  )
VALUES
  ($1, $2, $3, $4, $5) RETURNING id;