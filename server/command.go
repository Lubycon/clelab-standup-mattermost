package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
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

	// display the plugin version
	if strings.HasSuffix(command, "__version") {
		post := model.Post{
			ChannelId: args.ChannelId,
			UserId:    p.remindUserID,
			Message:   manifest.Version,
		}
		p.API.SendEphemeralPost(user.Id, &post)
		return &model.CommandResponse{}, nil
	}

	return &model.CommandResponse{}, nil
}
