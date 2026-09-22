package reports

import (
	"context"

	"github.com/rovioletta/standup-bot/internal/db"
	"github.com/rovioletta/standup-bot/internal/domain"
)

func (s *Service) SaveReport(ctx context.Context, report *domain.Report) error {
	_, error := s.queries.CreateReport(ctx, &db.CreateReportParams{
		UserID: report.UserID,
		TeamID: 0, // will be add later
		Report: report.Report,
	})

	return error
}