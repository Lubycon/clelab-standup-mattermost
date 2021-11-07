package main

import (
	"encoding/json"
	"strings"

	"github.com/mattermost/mattermost-plugin-starter-template/server/types"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"
)

const (
	CommandTrigger = "cletandup"

	UserAddedMessage                = "추가 성공!! 내일부터 스탠드업에서 만나요 :)"
	AlreadyRegisteredUserMessage    = "이미 등록된 유저입니다."
	AlreadyRegisteredChannelMessage = "이미 추가되어있는 채널입니다."
	ChannelDeletedMessage           = "삭제 성공 :)"
	ChannelAddedMessage             = "추가 성공 :)"
)

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

	if strings.HasSuffix(command, "apply") {
		channelID := args.ChannelId
		channelList := getChannels(p)

		for i, ch := range channelList {
			if ch.ID == channelID {
				for _, userID := range ch.Users {
					if userID == args.UserId {
						p.API.SendEphemeralPost(user.Id, &model.Post{
							ChannelId: args.ChannelId,
							UserId:    p.userID,
							Message:   AlreadyRegisteredUserMessage,
						})

						return &model.CommandResponse{}, nil
					}
				}

				channelList[i].Users = append(channelList[i].Users, args.UserId)

				channelJSON, _ := json.Marshal(channelList)
				_ = p.API.KVSet(ChannelListKey, channelJSON)

				p.API.SendEphemeralPost(user.Id, &model.Post{
					ChannelId: args.ChannelId,
					UserId:    p.userID,
					Message:   UserAddedMessage,
				})

				return &model.CommandResponse{}, nil
			}
		}
	}

	if strings.HasSuffix(command, "addChannel") {
		channelID := args.ChannelId
		channelList := getChannels(p)

		channel := types.Channel{ID: channelID, Users: []string{}}
		for _, ch := range channelList {
			if ch.ID == channel.ID {
				post := model.Post{
					ChannelId: args.ChannelId,
					UserId:    p.userID,
					Message:   AlreadyRegisteredChannelMessage,
				}
				p.API.SendEphemeralPost(user.Id, &post)
				return &model.CommandResponse{}, nil
			}
		}

		channelList = append(channelList, channel)

		channelJSON, err3 := json.Marshal(channelList)
		if err3 != nil {
			p.API.LogError(">>> [에러] Occurred error when Marshal : " + err3.Error())
			return &model.CommandResponse{}, nil
		}

		_ = p.API.KVSet(ChannelListKey, channelJSON)

		post := model.Post{
			ChannelId: args.ChannelId,
			UserId:    p.userID,
			Message:   ChannelAddedMessage,
		}
		p.API.SendEphemeralPost(user.Id, &post)
	}

	if strings.HasSuffix(command, "deleteChannel") {
		channelID := args.ChannelId
		channelList := getChannels(p)

		channel := types.Channel{ID: channelID}
		for i, ch := range channelList {
			if ch.ID == channel.ID {
				channelList = append(channelList[:i], channelList[i+1:]...)

				post := model.Post{
					ChannelId: args.ChannelId,
					UserId:    p.userID,
					Message:   ChannelDeletedMessage,
				}
				p.API.SendEphemeralPost(user.Id, &post)
				break
			}
		}

		channelJSON, _ := json.Marshal(channelList)
		_ = p.API.KVSet(ChannelListKey, channelJSON)
	}

	return &model.CommandResponse{}, nil
}
