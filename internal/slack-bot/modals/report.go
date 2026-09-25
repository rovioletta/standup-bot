package modals

import (
	"github.com/rovioletta/standup-bot/internal/slack-bot/constants"
	"github.com/slack-go/slack"
)

func ReportModal() slack.ModalViewRequest {
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
	return slack.ModalViewRequest{
		Type:       slack.VTModal,
		CallbackID: constants.CallbackDailyReportModalSubmit,
		Title:      slack.NewTextBlockObject("plain_text", "Daily Report", false, false),
		Submit:     slack.NewTextBlockObject("plain_text", "Submit Report", false, false),
		Close:      slack.NewTextBlockObject("plain_text", "Cancel", false, false),
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{reportBlock},
		},
	}
}
