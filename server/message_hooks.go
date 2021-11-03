package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"time"
)

func (p *Plugin) MessageHasBeenPosted(c *plugin.Context, post *model.Post) {
	channels := getChannels(p)
	user, _ := p.API.GetUser(post.UserId)

	nowTime := getNowTime()
	then := nowTime.Add(time.Duration(-4) * time.Hour)

	for _, channel := range channels {
		for _, userID := range channel.Users {
			if userID == post.UserId {
				dmChannel, _ := p.API.GetDirectChannel(post.UserId, p.userID)
				postList, _ := p.API.GetPostsSince(dmChannel.Id, then.UnixMilli())

				if len(postList.Posts) == 3 {
					p.PostBotDM(post.UserId, Question2)
				} else if len(postList.Posts) == 5 {
					p.PostBotDM(post.UserId, Question3)
				} else if len(postList.Posts) == 7 {
					p.PostBotDM(post.UserId, Question4)
				} else if len(postList.Posts) >= 8 {
					p.PostBotDM(post.UserId, StandUpCompleteMessage)
					postArray := postList.ToSlice()

					_, _ = p.API.CreatePost(&model.Post{
						ChannelId: channel.ID,
						UserId:    p.userID,
						Message: user.Username + "님의 오늘의 스탠드업! \n" +
							Question1 + "\n" +
							"-> " + postArray[2].Message + "\n" +
							Question2 + "\n" +
							"-> " + postArray[4].Message + "\n" +
							Question3 + "\n" +
							"-> " + postArray[6].Message + "\n" +
							Question4 + "\n" +
							"-> " + postArray[8].Message + "\n",
					})
				}
			}
		}
	}
}
