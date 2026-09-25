package commands

import (
	"log/slog"

	"github.com/rovioletta/standup-bot/internal/slack-bot/modals"
)

/*
/create_my_team -- open modal to create a team
*/
func (comhdl *CommandsHandler) createMyTeam(triggerID string) {
	_, err := comhdl.api.OpenView(triggerID, modals.CreateMyTeamModal())
	if err != nil {
		comhdl.logger.Error("Failed to open modal view", slog.String("error", err.Error()))
	}
}
