package commands

import (
	"log/slog"

	"github.com/rovioletta/standup-bot/internal/slack-bot/constants"
	"github.com/slack-go/slack"
)

/*
	/report -- builds and posts a message containing an interactive button
*/
func (comhdl *CommandsHandler) report(channelID string) {
	// Create interactive button element
	buttonBtn := slack.NewButtonBlockElement(
		constants.ActionOpenReportButton, // Unique ID used to identify the action in EventTypeInteractive
		"report_payload",                // Optional value passed with action
		slack.NewTextBlockObject("plain_text", "Fill Report", false, false),
	)

	buttonBtn.Style = slack.StylePrimary // Makes button green ("danger" for red, default for grey)

	// Create a section block holding the text and the button as an accessory
	sectionBlock := slack.NewSectionBlock(
		slack.NewTextBlockObject("mrkdwn", "Click the button below to start your report:\n", false, false),
		nil,
		nil,
	)

	actionsBlock := slack.NewActionBlock(
		"report_actions_block",
		buttonBtn,
	)

	// Send message containing the Block Kit layout
	_, _, err := comhdl.api.PostMessage(
		channelID,
		slack.MsgOptionBlocks(sectionBlock, actionsBlock),
	)
	if err != nil {
		comhdl.logger.Error("Failed to send message with button", slog.String("error", err.Error()))
	}
}
