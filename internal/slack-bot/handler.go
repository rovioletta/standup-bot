package slack_bot

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/rovioletta/standup-bot/internal/domain"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type ReportService interface {
	SaveReport(ctx context.Context, report *domain.Report) error
}

type SlackBot struct {
	api     *slack.Client
	client  *socketmode.Client
	logger  *slog.Logger
	reports ReportService
}

func New(logger *slog.Logger, reports ReportService) (*SlackBot, error) {
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
		api:     api,
		client:  client,
		logger:  logger,
		reports: reports,
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

			// 3. Handle Interactive Events (Button Clicks & Modal Submissions)
			case socketmode.EventTypeInteractive:
				interaction, ok := evt.Data.(slack.InteractionCallback)
				if !ok {
					continue
				}

				// Acknowledge interaction event immediately
				sb.client.Ack(*evt.Request)

				switch interaction.Type {
				// Handle button click action
				case slack.InteractionTypeBlockActions:
					for _, action := range interaction.ActionCallback.BlockActions {
						if action.ActionID == actionOpenReportButton {
							// Open modal using trigger_id from interaction payload
							sb.openReportSubmissionModal(interaction.TriggerID)
						}
					}

				// Handle modal form submission
				case slack.InteractionTypeViewSubmission:
					// Route based on CallbackID of the submitted modal
					switch interaction.View.CallbackID {

					case callbackDailyReportModalSubmit:

						go func(interaction slack.InteractionCallback) {
							err := sb.handleReportSubmissionModal(interaction)
							
							if err != nil {
								sb.logger.Error("Async report submission failed", "error", err)

								userID := interaction.User.ID

								msg := "Server Error. Please, copy your report and try later"
								_, _, sendErr := sb.api.PostMessage(userID, slack.MsgOptionText(msg, false))
								if sendErr != nil {
									sb.logger.Error("Failed to send error DM to user", "error", sendErr)
								}
							}
						}(interaction)

					default:
						sb.logger.Error("Unknown modal submission callback_id")
					}
				}
			}
		}
	}()

	sb.logger.Info("Bot started...")
	sb.client.Run()
}
