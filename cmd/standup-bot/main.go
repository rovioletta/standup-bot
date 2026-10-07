package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/rovioletta/standup-bot/internal/ai"
	"github.com/rovioletta/standup-bot/internal/cron"
	"github.com/rovioletta/standup-bot/internal/db"
	"github.com/rovioletta/standup-bot/internal/service/reports"
	"github.com/rovioletta/standup-bot/internal/service/teams"
	slack_bot "github.com/rovioletta/standup-bot/internal/slack-bot"
	"github.com/rovioletta/standup-bot/internal/slack-bot/commands"
	"github.com/rovioletta/standup-bot/internal/slack-bot/interactions"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func main() {
	// Define logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger.Info("Starting the bot...")

	// Read environment variables
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error: failed to load .env file", slog.String("error", err.Error()))
		return
	}

	// Postgres queries
	dbpool := initDB(logger)
	defer dbpool.Close()

	queries := db.New(dbpool)

	// AI tool
	ai := initAI(logger)

	// Create services
	reportSrv := reports.NewService(queries, ai)
	teamSrv := teams.NewService(queries)

	// Configure and start slack bot
	initBot(logger, reportSrv, teamSrv)
}

func initDB(logger *slog.Logger) *pgxpool.Pool {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		logger.Error("Database url is not provided")
		os.Exit(1)
	}

	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, url)
	if err != nil {
		logger.Error("Unable to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	if err := dbpool.Ping(ctx); err != nil {
		logger.Error("Unable to ping database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return dbpool
}

func initBot(logger *slog.Logger, reportSrv *reports.Service, teamSrv *teams.Service) {
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	if botToken == "" {
		logger.Error("SLACK_BOT_TOKEN is not provided")
		os.Exit(1)
	}

	appToken := os.Getenv("SLACK_APP_TOKEN")
	if appToken == "" {
		logger.Error("SLACK_APP_TOKEN is not provided")
		os.Exit(1)
	}

	api := slack.New(botToken, slack.OptionAppLevelToken(appToken))
	client := socketmode.New(api)

	cmdhdl := commands.NewCommandsHandler(client, api, logger)
	intmng := interactions.NewInteractionManager(client, api, logger, reportSrv, teamSrv)

	// Run cron jobs
	scheduler, err := cron.Run(logger, teamSrv, api)
	if err != nil {
		logger.Error("failed to start cron jobs", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer scheduler.Shutdown()

	slackBot := slack_bot.New(api, client, logger, cmdhdl, intmng)
	slackBot.StartEventsHandler()
}

func initAI(logger *slog.Logger) *ai.AI{
	client := openai.NewClient(
		option.WithBaseURL(os.Getenv("AI_API_URL")),
		option.WithAPIKey(os.Getenv("AI_API_KEY")),
	)

	ai, err := ai.New(client)
	if err != nil {
		logger.Error("failed to create ai client", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return ai
}
