package messages

import (
	"github.com/rovioletta/standup-bot/internal/slack-bot/constants"
	"github.com/slack-go/slack"
)

func SendReportButton(api *slack.Client, channelID, msg string) error {
	// Create interactive button element
	buttonBtn := slack.NewButtonBlockElement(
		constants.ActionOpenReportButton, // Unique ID used to identify the action in EventTypeInteractive
		"report_payload",                 // Optional value passed with action
		slack.NewTextBlockObject("plain_text", "Fill Report", false, false),
	)

	buttonBtn.Style = slack.StylePrimary // Makes button green ("danger" for red, default for grey)

	// Create a section block holding the text and the button as an accessory
	sectionBlock := slack.NewSectionBlock(
		slack.NewTextBlockObject("mrkdwn", msg + "\n", false, false),
		nil,
		nil,
	)

	actionsBlock := slack.NewActionBlock(
		"report_actions_block",
		buttonBtn,
	)

	// Send message containing the Block Kit layout
	_, _, err := api.PostMessage(
		channelID,
		slack.MsgOptionBlocks(sectionBlock, actionsBlock),
	)

	return err
}