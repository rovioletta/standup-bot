package slack_bot

import (
	"log/slog"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

type InteractionManager interface {
	Handle(evt socketmode.Event)
}

type CommandsHandler interface {
	Handle(evt socketmode.Event)
}

type SlackBot struct {
	api    *slack.Client
	client *socketmode.Client
	logger *slog.Logger
	cmdhdl CommandsHandler
	intmng InteractionManager
}

func New(
	api *slack.Client,
	client *socketmode.Client,
	logger *slog.Logger,
	cmdhdl CommandsHandler,
	intmng InteractionManager,
) *SlackBot {
	return &SlackBot{
		api:    api,
		client: client,
		logger: logger,
		cmdhdl: cmdhdl,
		intmng: intmng,
	}
}
