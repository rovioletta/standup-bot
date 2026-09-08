package slack_bot

import (
	"log/slog"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func (sb *SlackBot) directHandler(ev *slackevents.MessageEvent) {
	// Ignore bot's messages
	if ev.BotID != "" || ev.User == "" {
		return
	}

	userMessage := ev.Text
	sb.logger.Info("User texted in direct", slog.String("user", ev.User), slog.String("msg", userMessage))

	// Simplest handling
	if strings.ToLower(userMessage) == "hello" {
		sb.api.PostMessage(ev.Channel, slack.MsgOptionText("Hello!", false))
	} else {
		sb.api.PostMessage(ev.Channel, slack.MsgOptionText("I've recieved your message: "+userMessage, false))
	}
}
