package main

import (
	"github.com/mattermost/mattermost-plugin-api/cluster"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

const (
	TeamName        = "sss"
	SendChannelName = "town-square"
	TestUserID     = "ziba5knnofy9ucguaotfzwyz3h"
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
