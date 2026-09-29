package teams

import (
	"context"

	"github.com/rovioletta/standup-bot/internal/db"
	"github.com/rovioletta/standup-bot/internal/domain"
)

func (s *Service) SaveTeam(ctx context.Context, teamData *domain.Team) error {
	_, error := s.queries.CreateTeam(ctx, &db.CreateTeamParams{
		TeamName:         teamData.TeamName,
		ManagerID:        teamData.ManagerID,
		TeamMembers:      teamData.TeamMembers,
		ChannelID:        teamData.ChannelID,
		NotificationHour: teamData.NotificationHour,
	})

	return error
}
