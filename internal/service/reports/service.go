package reports

import (
	"context"

	"github.com/rovioletta/standup-bot/internal/db"
)

type Queries interface {
	CreateReport(ctx context.Context, arg *db.CreateReportParams) (uint64, error)
}

type AI interface{}

type Service struct {
	ai      AI
	queries Queries
}

func NewService(queries Queries, ai AI) *Service {
	return &Service{
		ai:      ai,
		queries: queries,
	}
}
