package slack_bot

import (
	"context"
	"log/slog"

	"github.com/rovioletta/standup-bot/internal/domain"
	"github.com/slack-go/slack"
)

// Opens the Modal View with text input and dropdown fields
func (sb *SlackBot) openReportSubmissionModal(triggerID string) {
	// Text input block: What was done
	reportInput := slack.NewPlainTextInputBlockElement(
		slack.NewTextBlockObject("plain_text", "Describe your completed tasks...", false, false),
		actionCreateReportManually,
	)
	reportInput.Multiline = true

	reportBlock := slack.NewInputBlock(
		blockReport,
		slack.NewTextBlockObject("plain_text", "What did you accomplish today?", false, false),
		nil,
		reportInput,
	)

	// Construct Modal Request View
	modalView := slack.ModalViewRequest{
		Type:       slack.VTModal,
		CallbackID: callbackDailyReportModalSubmit,
		Title:      slack.NewTextBlockObject("plain_text", "Daily Report", false, false),
		Submit:     slack.NewTextBlockObject("plain_text", "Submit Report", false, false),
		Close:      slack.NewTextBlockObject("plain_text", "Cancel", false, false),
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{reportBlock},
		},
	}

	// API call to render modal on user screen
	_, err := sb.api.OpenView(triggerID, modalView)
	if err != nil {
		sb.logger.Error("Failed to open modal view", slog.String("error", err.Error()))
	}
}

func (sb *SlackBot) handleReportSubmissionModal(interaction slack.InteractionCallback) (err error) {
	sb.logger.Info("Report", slog.Any("interaction", interaction.User.Name))

	if blockReportValue, ok := interaction.View.State.Values[blockReport]; ok {
		if report, ok := blockReportValue[actionCreateReportManually]; ok {
			sb.logger.Info("block report", slog.Any("report", report.Value))

			err = sb.reports.SaveReport(context.Background(), &domain.Report{
				UserID: interaction.User.ID,
				Report: report.Value,
			})
		}
	}

	return
}
