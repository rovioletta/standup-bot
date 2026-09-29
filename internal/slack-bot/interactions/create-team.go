package interactions

import (
	"context"
	"strconv"

	"github.com/rovioletta/standup-bot/internal/domain"
	"github.com/rovioletta/standup-bot/internal/slack-bot/constants"
	"github.com/slack-go/slack"
)

func (intmng *InteractionManager) handleCreateTeamSubmissionModal(interaction slack.InteractionCallback) (err error) {
	teamData := &domain.Team{}
	teamData.ManagerID = interaction.User.ID

	if blockValue, ok := interaction.View.State.Values[constants.BlockTeamName]; ok {
		if actionValue, ok := blockValue[constants.ActionMyTeamName]; ok {
			teamData.TeamName = actionValue.Value
		}
	}

	if blockValue, ok := interaction.View.State.Values[constants.BlockUsersSelect]; ok {
		if actionValue, ok := blockValue[constants.ActionSelectedUsers]; ok {
			teamData.TeamMembers = actionValue.SelectedUsers
		}
	}

	if blockValue, ok := interaction.View.State.Values[constants.BlockSelectedChannel]; ok {
		if actionValue, ok := blockValue[constants.ActionSelectedChannel]; ok {
			teamData.ChannelID = actionValue.SelectedChannel
		}
	}

	if blockValue, ok := interaction.View.State.Values[constants.BlockNotificationHour]; ok {
		if actionValue, ok := blockValue[constants.ActionSelectNotificationHour]; ok {
			timeStr := actionValue.SelectedOption.Value
			hour, _ := strconv.Atoi(timeStr[:2])
			teamData.NotificationHour = uint16(hour)
		}
	}

	err = intmng.teams.SaveTeam(context.Background(), teamData)
	return
}
