# Stand-up Slack Bot powered with AI

## About

A Slack bot that automates daily reporting, aggregates team updates, and generates structured summary reports for managers.


## Features

- [x] Command /report. Saving report to database with error handling.
- [x] Command /create-my-team. Saving manager settings.
- [x] Cron job for reminding team members to fill their reports.
- [ ] Open AI integration for reports aggregation.

---

## Slack Bot Commands (in progress)

| Command | Description | Access |
| :--- | :--- | :--- |
| **/report** | Triggers the report submission workflow. Sends a "Fill report" button that opens a modal window (allows selecting a date and team from a dropdown list). | All users |
| **/create_my_team** | Opens a modal window to create a team. Specifies team members, target chat for group reports and notification hour. The user who executes the command becomes the team manager. | All users |
| **/help** | Displays bot guidance, available commands, and usage instructions. | All users |
| **/get_my_report** | Fetches an individual report for a specific timeframe within a selected team. | All users |
| **/get_last** | Quick shortcut command to review the latest submitted report. | All users |
| **/get_team_report** | Manually retrieves an aggregated team report for a selected team managed by the user. | Team Manager |
| **/update_my_team** | Opens a modal window to manage team members (add/remove participants) and includes a dedicated button to delete the team. | Team Manager |
| **/my_team_status** | Checks the current report completion status to verify if all team members have submitted their reports. | Team Manager |

## Quick start

### 1. Environment Configuration

Copy and fill out the example environment file:

```bash
cp .env.example .env
```

### 2. Setup Bot Configuration
To enable slack bot work properly use slack-manifest.json. This file contains all settings and permissions.

### 3. Run with Docker Compose
Spin up the entire stack including the Go application, PostgreSQL database, and Adminer web UI. Database migrations located in ./migrations are automatically executed on startup:

```bash
docker-compose up -d --build
```

### 4. SQL Queries Generation
If you modify the database schemas or add new queries, run the generation script:

```bash
sqlc generate
```

See more by the link: [sqlc Documentation](https://docs.sqlc.dev/en/latest/)