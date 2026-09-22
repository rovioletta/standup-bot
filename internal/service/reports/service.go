package reports

import (
	"context"

	"github.com/rovioletta/standup-bot/internal/db"
)

type Queries interface {
	CreateReport(ctx context.Context, arg *db.CreateReportParams) (uint64, error)
}

type Service struct {
	queries Queries
}

func NewService(queries Queries) *Service {
	return &Service{
		queries: queries,
	}
}
