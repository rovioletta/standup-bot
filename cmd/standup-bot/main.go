package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rovioletta/standup-bot/internal/db"
	"github.com/rovioletta/standup-bot/internal/service/reports"
	slack_bot "github.com/rovioletta/standup-bot/internal/slack-bot"
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

	// Create Report Service
	reportService := reports.NewService(queries)

	// Configure and start slack bot
	slackBot, err := slack_bot.New(logger, reportService)
	if err != nil {
		logger.Error("Error: failed to create slack bot", slog.String("error", err.Error()))
		return
	}

	slackBot.StartEventsHandler()
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
