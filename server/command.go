package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"
)

const CommandTrigger = "cletandup"

func (p *Plugin) registerCommand() error {
	if err := p.API.RegisterCommand(&model.Command{
		Trigger:          CommandTrigger,
		AutoComplete:     true,
		AutoCompleteHint: "hint",
		AutoCompleteDesc: "desc",
	}); err != nil {
		return errors.Wrap(err, "failed to register command")
	}

	return nil
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	user, uErr := p.API.GetUser(args.UserId)
	if uErr != nil {
		return &model.CommandResponse{}, uErr
	}

	command := strings.Trim(args.Command, " ")

	if strings.HasSuffix(command, "__version") {
		post := model.Post{
			ChannelId: args.ChannelId,
			UserId:    p.userID,
			Message:   manifest.Version,
		}
		p.API.SendEphemeralPost(user.Id, &post)
		return &model.CommandResponse{}, nil
	}

	if strings.HasSuffix(command, "__id") {
		post := model.Post{
			ChannelId: args.ChannelId,
			UserId:    p.userID,
			Message:   p.userID,
		}
		p.API.SendEphemeralPost(user.Id, &post)
		return &model.CommandResponse{}, nil
	}

	if strings.HasSuffix(command, "send") {
		message := "안냐쎄여 저는 클랩의 스탠드업을 책임져줄 로봇이에요!"
		p.PostBotDM("ziba5knnofy9ucguaotfzwyz3h", message) //FIXME: 여기 바꿔야됨!
	}

	return &model.CommandResponse{}, nil
}
