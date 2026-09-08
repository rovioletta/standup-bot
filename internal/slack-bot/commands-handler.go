package slack_bot

import (
	"fmt"
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

	// Response
	reply := fmt.Sprintf("Command `%s` with params %s were executed", cmd.Command, cmd.Text)
	sb.api.PostMessage(cmd.ChannelID, slack.MsgOptionText(reply, false))
}