package main

import (
	"encoding/json"
	"strings"

	"github.com/mattermost/mattermost-plugin-starter-template/server/types"
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
		channelID := args.ChannelId
		channelListData, err := p.API.KVGet(ChannelListKey)
		if err != nil {
			p.API.LogError(">>> [에러] Occurred error when KVGet : " + err.Error())
			return &model.CommandResponse{}, err
		}

		channelList := types.ChannelList{}
		err2 := json.Unmarshal(channelListData, &channelList)
		if err2 != nil {
			p.API.LogError(">>> [에러] Occurred error when Unmarshal : " + err2.Error())
		}

		channel := types.Channel{ID: channelID}
		channelList = append(channelList, channel)

		channelJSON, err3 := json.Marshal(channelList)
		if err3 != nil {
			p.API.LogError(">>> [에러] Occurred error when Marshal : " + err3.Error())
			return &model.CommandResponse{}, nil
		}

		err = p.API.KVSet(ChannelListKey, channelJSON)
		if err != nil {
			p.API.LogError(">>> [에러] Occurred error when KVSet : " + err.Error())
			return &model.CommandResponse{}, err
		}

		post := model.Post{
			ChannelId: args.ChannelId,
			UserId:    p.userID,
			Message:   "추가완료!",
		}
		p.API.SendEphemeralPost(user.Id, &post)
	}

	return &model.CommandResponse{}, nil
}
