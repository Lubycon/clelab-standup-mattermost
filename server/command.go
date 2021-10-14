package main

import (
	"encoding/json"
	"github.com/mattermost/mattermost-plugin-starter-template/server/type"
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
		message := StandUpMessage
		p.PostBotDM(TestUserID, message) //FIXME: 여기 바꿔야됨!
	}

	if strings.HasSuffix(command, "addChannel") {
		channelId := args.ChannelId
		channelListData, err := p.API.KVGet(ChannelListKey)
		if err != nil {
			return &model.CommandResponse{}, err
		}

		channelList := _type.ChannelList{}
		err2 := json.Unmarshal(channelListData, &channelList)
		if err2 != nil {
			return &model.CommandResponse{}, nil
		}

		channel := _type.Channel{ID: channelId}
		channelList = append(channelList, channel)

		channelJson, err3 := json.Marshal(channelList)
		if err3 != nil {
			return &model.CommandResponse{}, nil
		}

		err = p.API.KVSet(ChannelListKey, channelJson)
		if err != nil {
			return &model.CommandResponse{}, err
		}
	}

	return &model.CommandResponse{}, nil
}
