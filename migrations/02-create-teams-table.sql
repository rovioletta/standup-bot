CREATE TABLE IF NOT EXISTS teams (
  id SERIAL PRIMARY KEY,
  team_name TEXT NOT NULL,
  manager_id TEXT NOT NULL,
  team_members TEXT[] NOT NULL,
  channel_id TEXT NOT NULL,
  notification_hour smallint NOT NULL
);