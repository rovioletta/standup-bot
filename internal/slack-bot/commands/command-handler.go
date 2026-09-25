package commands

import (
	"log/slog"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

type CommandsHandler struct {
	client *socketmode.Client
	api    *slack.Client
	logger *slog.Logger
}

func NewCommandsHandler(client *socketmode.Client, api *slack.Client, logger *slog.Logger) *CommandsHandler {
	return &CommandsHandler{
		client: client,
		api:    api,
		logger: logger,
	}
}

func (comhdl *CommandsHandler) Handle(evt socketmode.Event) {
	cmd, ok := evt.Data.(slack.SlashCommand)
	if !ok {
		return
	}
	comhdl.client.Ack(*evt.Request) // Mark as readed

	comhdl.logger.Info("User executed command with params",
		slog.String("user", cmd.UserID),
		slog.String("command", cmd.Command),
		slog.String("params", cmd.Text),
	)

	switch cmd.Command {
	case "/report":
		comhdl.report(cmd.ChannelID)
	case "/create_my_team":
		comhdl.createMyTeam(cmd.TriggerID)
	}

}
