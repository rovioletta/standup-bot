package slack_bot

import (
	"log/slog"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func (sb *SlackBot) commandsHandler(evt socketmode.Event) {
	cmd, ok := evt.Data.(slack.SlashCommand)
	if !ok {
		return
	}
	sb.client.Ack(*evt.Request) // Mark as readed

	sb.logger.Info("User executed command with params",
		slog.String("user", cmd.UserID),
		slog.String("command", cmd.Command),
		slog.String("params", cmd.Text),
	)

	switch cmd.Command {
	case "/report":
		sb.sendReportButton(cmd.ChannelID)
	}
}

// Builds and posts a message containing an interactive button
func (sb *SlackBot) sendReportButton(channelID string) {
	// Create interactive button element
	// Parameters: action_id, value, text_object
	buttonBtn := slack.NewButtonBlockElement(
		actionOpenReportButton, // Unique ID used to identify the action in EventTypeInteractive
		"report_payload",       // Optional value passed with action
		slack.NewTextBlockObject("plain_text", "Fill Report", false, false),
	)

	buttonBtn.Style = slack.StylePrimary // Makes button green ("danger" for red, default for grey)

	// Create a section block holding the text and the button as an accessory
	sectionBlock := slack.NewSectionBlock(
		slack.NewTextBlockObject("mrkdwn", "Click the button below to start your report:\n", false, false),
		nil,
		nil,
		//slack.NewAccessory(buttonBtn),
	)

	actionsBlock := slack.NewActionBlock(
		"report_actions_block",
		buttonBtn,
	)

	// Send message containing the Block Kit layout
	_, _, err := sb.api.PostMessage(
		channelID,
		slack.MsgOptionBlocks(sectionBlock, actionsBlock),
	)
	if err != nil {
		sb.logger.Error("Failed to send message with button", slog.String("error", err.Error()))
	}
}
