package cron

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-co-op/gocron/v2"
	"github.com/slack-go/slack"
)

type TeamService interface {
	GetMembersToNotify(ctx context.Context) ([]string, error)
}

type Scheduler struct {
	goSch    gocron.Scheduler
	logger   *slog.Logger
	slackApi *slack.Client
	teams    TeamService
}

func Run(logger *slog.Logger, teams TeamService, slackApi *slack.Client) (*Scheduler, error) {
	// create a scheduler
	goSch, err := gocron.NewScheduler()
	if err != nil {
		return nil, fmt.Errorf("failed to create scheduler: %w", err)
	}

	sch := &Scheduler{
		goSch:    goSch,
		logger:   logger,
		slackApi: slackApi,
		teams:    teams,
	}

	// add a job to the scheduler
	_, err = goSch.NewJob(
		gocron.CronJob("* * * * *", false),
		gocron.NewTask(sch.notifyTeams),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to add a job: %w", err)
	}

	// start the scheduler
	goSch.Start()

	return sch, nil
}

func (sch *Scheduler) Shutdown() error {
	sch.logger.Info("Stop cron jobs...")
	err := sch.goSch.Shutdown()
	if err != nil {
		return fmt.Errorf("error while shutdown scheduler: %w", err)
	}

	return nil
}
