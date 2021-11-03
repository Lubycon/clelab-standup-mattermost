package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func (p *Plugin) MessageHasBeenPosted(c *plugin.Context, post *model.Post) {
	channels := getChannels(p)
	user, _ := p.API.GetUser(post.UserId)

	for _, channel := range channels {
		for _, userID := range channel.Users {
			if userID == post.UserId {
				_, _ = p.API.CreatePost(&model.Post{
					ChannelId: channel.ID,
					UserId:    p.userID,
					Message:   user.Username + "님, \n" + post.Message,
				})
			}
		}
	}
}
