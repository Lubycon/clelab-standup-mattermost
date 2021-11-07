package main

import (
	"strings"
	"time"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/pkg/errors"
)

func SendNotification(p *Plugin, nowTime time.Time) error {
	channelList := getChannels(p)

	for _, channel := range channelList {
		_, err := p.API.CreatePost(&model.Post{
			ChannelId: channel.ID,
			UserId:    p.userID,
			Message:   nowTime.Format("2006년 01월 02일 오늘의 스탠드업!"),
		})
		if err != nil {
			return err
		}

		ids := channel.Users
		for _, id := range ids {
			p.PostBotDM(id, StandUpMessage)
			p.PostBotDM(id, Question1)
		}
	}
	return nil
}

func SendReminder(p *Plugin, nowTime time.Time) error {
	channelList := getChannels(p)
	then := nowTime.Add(time.Duration(-4) * time.Hour)

	for _, channel := range channelList {
		ids := channel.Users

		for _, id := range ids {
			dmChannel, appError := p.API.GetDirectChannel(id, p.userID)
			if appError != nil {
				return appError
			}

			time.Now().UnixMilli()

			postList, postErr := p.API.GetPostsSince(dmChannel.Id, then.UnixMilli())
			if postErr != nil {
				return appError
			}

			if len(postList.Posts) == 2 {
				p.PostBotDM(id, StandUpRemindMessage)
			}
		}
	}
	return nil
}

func (p *Plugin) PostBotDM(userID string, message string) {
	p.createBotPostDM(&model.Post{
		UserId:  p.userID,
		Message: message,
	}, userID)
}

func (p *Plugin) createBotPostDM(post *model.Post, userID string) {
	channel, appError := p.API.GetDirectChannel(userID, p.userID)

	if appError != nil {
		p.API.LogError(">>> [에러] Unable to get direct channel for bot err: " + appError.Error())
		return
	}
	if channel == nil {
		p.API.LogError(">>> [에러] Could not get direct channel for bot and user_id: " + appError.Error())
		return
	}

	post.ChannelId = channel.Id
	_, appError = p.API.CreatePost(post)

	if appError != nil {
		p.API.LogError(">>> [에러] Unable to create bot post DM err: " + appError.Error())
	}
}

func (p *Plugin) ReplyPostBot(postID, message, todo string) error {
	if postID == "" {
		return errors.New(">>> [에러] post ID not defined")
	}

	post, appErr := p.API.GetPost(postID)
	if appErr != nil {
		return appErr
	}
	rootID := post.Id
	if post.RootId != "" {
		rootID = post.RootId
	}

	quotedTodo := "\n> " + strings.Join(strings.Split(todo, "\n"), "\n> ")
	_, appErr = p.API.CreatePost(&model.Post{
		UserId:    p.userID,
		ChannelId: post.ChannelId,
		Message:   message + quotedTodo,
		RootId:    rootID,
	})

	if appErr != nil {
		return appErr
	}

	return nil
}
