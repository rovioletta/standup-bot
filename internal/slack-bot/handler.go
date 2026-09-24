package slack_bot

import (
	"github.com/slack-go/slack/socketmode"
)

func (sb *SlackBot) StartEventsHandler() {
	go func() {
		for evt := range sb.client.Events {
			switch evt.Type {

			// 1. Handling slesh commands (for example /hello)
			case socketmode.EventTypeSlashCommand:
				sb.cmdhdl.Handle(evt)

			// 2. Handle Interactive Events (Button Clicks & Modal Submissions)
			case socketmode.EventTypeInteractive:
				sb.intmng.Handle(evt)
			}
		}
	}()

	sb.logger.Info("Bot started...")
	sb.client.Run()
}
