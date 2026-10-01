-- name: GetMembersToNotify :one
SELECT
  ARRAY_AGG(DISTINCT member) :: text[] AS team_members
FROM
  teams,
  UNNEST(team_members) AS member
WHERE
  notification_hour = @notification_hour :: smallint;