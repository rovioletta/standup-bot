package modals

import (
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

	// Construct Modal Request View
	return slack.ModalViewRequest{
		Type:       slack.VTModal,
		CallbackID: constants.CallbackDailyReportModalSubmit,
		Title:      slack.NewTextBlockObject("plain_text", "New Team", false, false),
		Submit:     slack.NewTextBlockObject("plain_text", "Submit", false, false),
		Close:      slack.NewTextBlockObject("plain_text", "Cancel", false, false),
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{teamNameBlock, usersSelectBlock, channelSelectBlock},
		},
	}
}
