package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
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

	// Configure and start slack bot
	slackBot, err := slack_bot.New(logger)
	if err != nil {
		logger.Error("Error: failed to create slack bot", slog.String("error", err.Error()))
		return
	}

	slackBot.StartEventsHandler()
}
