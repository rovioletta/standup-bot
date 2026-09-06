package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func main() {
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	api := slack.New(botToken, slack.OptionAppLevelToken(appToken))
	client := socketmode.New(api)

	go func() {
		for evt := range client.Events {
			switch evt.Type {

			// 1. Обработка обычных событий (ЛС и Упоминания)
			case socketmode.EventTypeEventsAPI:
				eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
				if !ok {
					continue
				}
				client.Ack(*evt.Request) // Подтверждаем получение

				if eventsAPIEvent.Type == slackevents.CallbackEvent {
					switch ev := eventsAPIEvent.InnerEvent.Data.(type) {

					// Пользователь упомянул бота в канале: @BotText
					case *slackevents.AppMentionEvent:
						userMessage := ev.Text // Читаем текст сообщения
						log.Printf("Пользователь %s упомянул бота: %s", ev.User, userMessage)

						reply := fmt.Sprintf("Вы упомянули меня с текстом: *%s*", userMessage)
						api.PostMessage(ev.Channel, slack.MsgOptionText(reply, false))

					// Пользователь написал личное сообщение боту (ЛС)
					case *slackevents.MessageEvent:
						// Игнорируем сообщения от самого бота (чтобы не уйти в бесконечный цикл)
						if ev.BotID != "" || ev.User == "" {
							continue
						}

						userMessage := ev.Text // Читаем текст сообщения
						log.Printf("Пользователь %s написал в ЛС: %s", ev.User, userMessage)

						// Простейшая обработка текста
						if strings.ToLower(userMessage) == "привет" {
							api.PostMessage(ev.Channel, slack.MsgOptionText("Привет! Чем могу помочь?", false))
						} else {
							api.PostMessage(ev.Channel, slack.MsgOptionText("Я получил твое сообщение: "+userMessage, false))
						}
					}
				}

			// 2. Handling slesh commands (for example /hello)
			case socketmode.EventTypeSlashCommand:
				cmd, ok := evt.Data.(slack.SlashCommand)
				if !ok {
					continue
				}
				client.Ack(*evt.Request) // Mark as readed

				log.Printf("Пользователь %s вызвал команду %s с аргументами: %s", cmd.UserID, cmd.Command, cmd.Text)

				// Отвечаем на слэш-команду
				reply := fmt.Sprintf("Команда `%s` выполнена с параметрами: %s", cmd.Command, cmd.Text)
				api.PostMessage(cmd.ChannelID, slack.MsgOptionText(reply, false))
			}
		}
	}()

	log.Println("Бот запущен и слушает сообщения...")
	client.Run()
}