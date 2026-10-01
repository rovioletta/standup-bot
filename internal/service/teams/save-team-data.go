package teams

import (
	"context"
	"fmt"

	"github.com/rovioletta/standup-bot/internal/db"
	"github.com/rovioletta/standup-bot/internal/domain"
)

func (s *Service) SaveTeam(ctx context.Context, teamData *domain.Team) error {
	_, err := s.queries.CreateTeam(ctx, &db.CreateTeamParams{
		TeamName:         teamData.TeamName,
		ManagerID:        teamData.ManagerID,
		TeamMembers:      teamData.TeamMembers,
		ChannelID:        teamData.ChannelID,
		NotificationHour: teamData.NotificationHour,
	})
	if err != nil {
		return fmt.Errorf("failed to save team to db: %w", err)
	}

	return nil
}
