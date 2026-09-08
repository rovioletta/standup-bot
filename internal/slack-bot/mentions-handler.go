package slack_bot

import (
	"fmt"
	"log/slog"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func (sb *SlackBot) mentionsHandler(ev *slackevents.AppMentionEvent) {
	userMessage := ev.Text
	sb.logger.Info("User mentioned bot", slog.String("user", ev.User), slog.String("msg", userMessage))

	reply := fmt.Sprintf("You mentioned me with text: *%s*", userMessage)
	sb.api.PostMessage(ev.Channel, slack.MsgOptionText(reply, false))
}
