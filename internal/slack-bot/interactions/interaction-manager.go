package interactions

import (
	"context"
	"log/slog"

	"github.com/rovioletta/standup-bot/internal/domain"
	"github.com/rovioletta/standup-bot/internal/slack-bot/constants"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

type ReportService interface {
	SaveReport(ctx context.Context, report *domain.Report) error
}

type InteractionManager struct {
	client  *socketmode.Client
	api     *slack.Client
	logger  *slog.Logger
	reports ReportService
}

func NewInteractionManager(client *socketmode.Client, api *slack.Client, logger *slog.Logger, reports ReportService) *InteractionManager {
	return &InteractionManager{
		client:  client,
		api:     api,
		logger:  logger,
		reports: reports,
	}
}

func (intmng *InteractionManager) Handle(evt socketmode.Event) {
	interaction, ok := evt.Data.(slack.InteractionCallback)
	if !ok {
		return
	}

	// Acknowledge interaction event immediately
	intmng.client.Ack(*evt.Request)

	switch interaction.Type {
	// Handle button click action
	case slack.InteractionTypeBlockActions:
		for _, action := range interaction.ActionCallback.BlockActions {
			if action.ActionID == constants.ActionOpenReportButton {
				// Open modal using trigger_id from interaction payload
				intmng.openReportSubmissionModal(interaction.TriggerID)
			}
		}

	// Handle modal form submission
	case slack.InteractionTypeViewSubmission:
		// Route based on CallbackID of the submitted modal
		switch interaction.View.CallbackID {

		case constants.CallbackDailyReportModalSubmit:

			go func(interaction slack.InteractionCallback) {
				report, err := intmng.handleReportSubmissionModal(interaction)

				userID := interaction.User.ID

				if err != nil {
					intmng.logger.Error("Async report submission failed", "error", err)

					msg := "Server Error. Please copy your report and try later\n\n" + report
					_, _, sendErr := intmng.api.PostMessage(userID, slack.MsgOptionText(msg, false))
					if sendErr != nil {
						intmng.logger.Error("Failed to send error DM to user", "error", sendErr)
					}
				}

				intmng.api.PostMessage(userID, slack.MsgOptionText("Your report was successfully saved!", false))

			}(interaction)

		default:
			intmng.logger.Error("Unknown modal submission callback_id")
		}
	}
}
