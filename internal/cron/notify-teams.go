package cron

import (
	"context"
	"log/slog"

	"github.com/rovioletta/standup-bot/internal/slack-bot/messages"
)

func (sch *Scheduler) notifyTeams() {
	// get all teams by notification_hour
	userIDs, err := sch.teams.GetMembersToNotify(context.Background())
	if err != nil {
		sch.logger.Error("cronjob: sch.teams.GetMembersToNotify", slog.String("error", err.Error()))
	}

	// send messages to DMs
	for _, userID := range userIDs {
		sch.sendNotificationToUser(userID)
	}
}

func (sch *Scheduler) sendNotificationToUser(userID string) {
	err := messages.SendReportButton(sch.slackApi, userID, "It's time to fill your report!")
	if err != nil {
		sch.logger.Error("cronjob: messages.SendReportButton", slog.String("error", err.Error()))
	}
}