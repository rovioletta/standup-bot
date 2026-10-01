package commands

import (
	"log/slog"

	"github.com/rovioletta/standup-bot/internal/slack-bot/messages"
)

/*
/report -- builds and posts a message containing an interactive button
*/
func (comhdl *CommandsHandler) report(channelID string) {
	err := messages.SendReportButton(comhdl.api, channelID, "Click the button below to start your report:")
	if err != nil {
		comhdl.logger.Error("Failed to send message with button", slog.String("error", err.Error()))
	}
}
