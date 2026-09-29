package teams

import (
	"context"

	"github.com/rovioletta/standup-bot/internal/db"
)

type Queries interface {
	CreateTeam(ctx context.Context, arg *db.CreateTeamParams) (uint64, error)
}

type Service struct {
	queries Queries
}

func NewService(queries Queries) *Service {
	return &Service{
		queries: queries,
	}
}
