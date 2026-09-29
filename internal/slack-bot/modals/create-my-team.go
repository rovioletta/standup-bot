package modals

import (
	"fmt"

	"github.com/rovioletta/standup-bot/internal/slack-bot/constants"
	"github.com/slack-go/slack"
)

func CreateMyTeamModal() slack.ModalViewRequest {
	teamNameElement := slack.NewPlainTextInputBlockElement(
		slack.NewTextBlockObject("plain_text", "Create your team...", false, false),
		constants.ActionMyTeamName,
	)

	teamNameBlock := slack.NewInputBlock(
		constants.BlockTeamName,
		slack.NewTextBlockObject(slack.PlainTextType, "Name your team:", false, false),
		nil,
		teamNameElement,
	)

	usersSelectElement := slack.NewOptionsMultiSelectBlockElement(
		slack.MultiOptTypeUser,
		slack.NewTextBlockObject(slack.PlainTextType, "Start to input nick or name...", false, false),
		constants.ActionSelectedUsers,
	)

	usersSelectBlock := slack.NewInputBlock(
		constants.BlockUsersSelect,
		slack.NewTextBlockObject(slack.PlainTextType, "Choose teammates:", false, false),
		nil,
		usersSelectElement,
	)

	channelSelectElement := slack.NewOptionsMultiSelectBlockElement(
		slack.OptTypeChannels,
		slack.NewTextBlockObject(slack.PlainTextType, "Choose one channel...", false, false),
		constants.ActionSelectedChannel,
	)

	channelSelectBlock := slack.NewInputBlock(
		constants.BlockSelectedChannel,
		slack.NewTextBlockObject(slack.PlainTextType, "Report Channel", false, false),
		slack.NewTextBlockObject(slack.PlainTextType, "If left empty, reports will be sent to your Direct Messages.", false, false),
		channelSelectElement,
	)
	channelSelectBlock.Optional = true

	var options []*slack.OptionBlockObject
	for hour := range 24 {
		timeStr := fmt.Sprintf("%02d:00", hour)
		
		option := slack.NewOptionBlockObject(
			timeStr,
			slack.NewTextBlockObject(slack.PlainTextType, timeStr, false, false),
			nil,
		)
		options = append(options, option)
	}

	timeSelectElement := slack.NewOptionsSelectBlockElement(
		slack.OptTypeStatic,
		slack.NewTextBlockObject(slack.PlainTextType, "Select hour...", false, false),
		constants.ActionSelectNotificationHour,
		options...,
	)

	timeBlock := slack.NewInputBlock(
		constants.BlockNotificationHour,
		slack.NewTextBlockObject(slack.PlainTextType, "Notification hour", false, false),
		slack.NewTextBlockObject(slack.PlainTextType, "Reports are sent at the top of the hour.", false, false),
		timeSelectElement,
	)

	return slack.ModalViewRequest{
		Type:       slack.VTModal,
		CallbackID: constants.CallbackCreateTeamModalSubmit,
		Title:      slack.NewTextBlockObject("plain_text", "New Team", false, false),
		Submit:     slack.NewTextBlockObject("plain_text", "Submit", false, false),
		Close:      slack.NewTextBlockObject("plain_text", "Cancel", false, false),
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{teamNameBlock, usersSelectBlock, channelSelectBlock, timeBlock},
		},
	}
}
