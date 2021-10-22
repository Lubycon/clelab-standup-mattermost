package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func (p *Plugin) MessageHasBeenPosted(c *plugin.Context, post *model.Post) {
	channel, err := p.API.GetDirectChannel(post.UserId, p.userID)
	if err != nil {
		p.API.LogError(
			"Failed to query direct channel",
			"user_id", post.UserId,
			"error", err.Error(),
		)
		return
	}

	if post.ChannelId == channel.Id {
		team, err := p.API.GetTeamByName(TeamName)
		if err != nil {
			p.API.LogError(
				"Failed to query user",
				"user_id", post.UserId,
				"error", err.Error(),
			)
			return
		}

		targetChannel, err := p.API.GetChannelByName(team.Id, SendChannelName, false)
		if err != nil {
			p.API.LogError(
				"Failed to query channel",
				"channel_id", targetChannel.Id,
				"error", err.Error(),
			)
			return
		}

		_, pErr := p.API.CreatePost(&model.Post{
			ChannelId: targetChannel.Id,
			UserId:    p.userID,
			Message:   post.Message,
		})
		if pErr != nil {
			p.API.LogError(
				"Failed to CreatePost : " + pErr.Error(),
			)
			return
		}
	}
}
