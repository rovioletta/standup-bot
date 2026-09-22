package slack_bot

import (
	"fmt"
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

func (sb *SlackBot) mentionsHandler(ev *slackevents.AppMentionEvent) {
	userMessage := ev.Text
	sb.logger.Info("User mentioned bot", slog.String("user", ev.User), slog.String("msg", userMessage))

	reply := fmt.Sprintf("You mentioned me with text: *%s*", userMessage)
	sb.api.PostMessage(ev.Channel, slack.MsgOptionText(reply, false))
}
