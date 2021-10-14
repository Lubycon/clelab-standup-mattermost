package main

import (
	"github.com/mattermost/mattermost-plugin-api/cluster"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)


const (
	ChannelListKey  = "channelList"
	TeamName        = "lubycon"
	SendChannelName = "town-square" // 여기가, 유저가 스탠덥 메세지를 보내면 결과를 보내주는 채널 > 배열로 관리하기.
	TestUserID      = "ziba5knnofy9ucguaotfzwyz3h" // 이거는 스탠드업에 참여할 멤버 리스트.
	StandUpMessage  = `오늘의 스탠드업을 팀원들과 공유해주세요!
1. 어제는 어떤 일을 하셨나요?
2. 만약, 어제 계획했던 일을 하지 못했다면 그 이유는 무엇인가요?
3. 오늘 해야할 일은 무엇인가요?
4. 공유사항이 있다면 자유롭게 적어주세요!`
)

type Plugin struct {
	plugin.MattermostPlugin

	ServerConfig *model.Config

	userID string

	job *cluster.Job
}
