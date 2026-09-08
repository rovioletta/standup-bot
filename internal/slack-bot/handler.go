package slack_bot

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type SlackBot struct {
	api    *slack.Client
	client *socketmode.Client
	logger *slog.Logger
}

func New(logger *slog.Logger) (*SlackBot, error) {
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	if botToken == "" {
		return nil, fmt.Errorf("error: SLACK_BOT_TOKEN is not provided")
	}

	appToken := os.Getenv("SLACK_APP_TOKEN")
	if appToken == "" {
		return nil, fmt.Errorf("error: SLACK_APP_TOKEN is not provided")
	}

	api := slack.New(botToken, slack.OptionAppLevelToken(appToken))
	client := socketmode.New(api)

	return &SlackBot{
		api:    api,
		client: client,
		logger: logger,
	}, nil
}

func (sb *SlackBot) StartEventsHandler() {
	go func() {
		for evt := range sb.client.Events {
			switch evt.Type {
			// 1. Handling messages and mentions
			case socketmode.EventTypeEventsAPI:
				eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
				if !ok {
					continue
				}
				sb.client.Ack(*evt.Request)

				if eventsAPIEvent.Type == slackevents.CallbackEvent {
					switch ev := eventsAPIEvent.InnerEvent.Data.(type) {

					case *slackevents.AppMentionEvent:
						sb.mentionsHandler(ev)

					case *slackevents.MessageEvent:
						sb.directHandler(ev)
					}
				}

			// 2. Handling slesh commands (for example /hello)
			case socketmode.EventTypeSlashCommand:
				sb.commandsHandler(evt)
			}
		}
	}()

	sb.logger.Info("Bot started...")
	sb.client.Run()
}
