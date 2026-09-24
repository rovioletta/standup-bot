package interactions
import (
	"context"
	"log/slog"

	"github.com/rovioletta/standup-bot/internal/domain"
	"github.com/rovioletta/standup-bot/internal/slack-bot/constants"
	"github.com/slack-go/slack"
)

// Opens the Modal View with text input and dropdown fields
func (intmng *InteractionManager) openReportSubmissionModal(triggerID string) {
	// Text input block: What was done
	reportInput := slack.NewPlainTextInputBlockElement(
		slack.NewTextBlockObject("plain_text", "Describe your completed tasks...", false, false),
		constants.ActionCreateReportManually,
	)
	reportInput.Multiline = true

	reportBlock := slack.NewInputBlock(
		constants.BlockReport,
		slack.NewTextBlockObject("plain_text", "What did you accomplish today?", false, false),
		nil,
		reportInput,
	)

	// Construct Modal Request View
	modalView := slack.ModalViewRequest{
		Type:       slack.VTModal,
		CallbackID: constants.CallbackDailyReportModalSubmit,
		Title:      slack.NewTextBlockObject("plain_text", "Daily Report", false, false),
		Submit:     slack.NewTextBlockObject("plain_text", "Submit Report", false, false),
		Close:      slack.NewTextBlockObject("plain_text", "Cancel", false, false),
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{reportBlock},
		},
	}

	// API call to render modal on user screen
	_, err := intmng.api.OpenView(triggerID, modalView)
	if err != nil {
		intmng.logger.Error("Failed to open modal view", slog.String("error", err.Error()))
	}
}

func (intmng *InteractionManager) handleReportSubmissionModal(interaction slack.InteractionCallback) (result string, err error) {
	if blockReportValue, ok := interaction.View.State.Values[constants.BlockReport]; ok {
		if report, ok := blockReportValue[constants.ActionCreateReportManually]; ok {
			intmng.logger.Info("block report", slog.Any("report", report.Value))

			err = intmng.reports.SaveReport(context.Background(), &domain.Report{
				UserID: interaction.User.ID,
				Report: report.Value,
			})

			result = report.Value
		}
	}

	return
}
