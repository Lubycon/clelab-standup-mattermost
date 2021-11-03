package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"

	"time"
)

const (
	StandUpMessage         = "오늘의 스탠드업을 팀원들과 공유해주세요!"
	StandUpCompleteMessage = "스탠드업 작성 완료! 💪"

	Question1 = "1. 어제는 어떤 일을 하셨나요?"
	Question2 = "2. 만약, 어제 계획했던 일을 하지 못했다면 그 이유는 무엇인가요?"
	Question3 = "3. 오늘 해야할 일은 무엇인가요?"
	Question4 = "4. 공유사항이 있다면 자유롭게 적어주세요!"
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

				switch {
				case len(postList.Posts) >= 3:
					p.PostBotDM(post.UserId, Question2)
				case len(postList.Posts) >= 5:
					p.PostBotDM(post.UserId, Question3)
				case len(postList.Posts) >= 7:
					p.PostBotDM(post.UserId, Question4)
				case len(postList.Posts) >= 8:
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
				default:
				}
			}
		}
	}
}
