package main

import (
	"time"

	"github.com/mattermost/mattermost-plugin-api/cluster"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/pkg/errors"
)

func (p *Plugin) OnActivate() error {
	p.ServerConfig = p.API.GetConfig()
	p.userID = p.settingBotInfo()
	p.job = p.settingScheduler()

	if p.registerCommand() != nil {
		return errors.Wrap(p.registerCommand(), "failed to register command")
	}

	return nil
}

func (p *Plugin) settingBotInfo() string {
	botUserID, err := p.Helpers.EnsureBot(&model.Bot{

		Username:    "cletandup",
		DisplayName: "Cletandup",
		Description: "Hello! I'm standup bot for clelab.",
	})

	if err != nil {
		p.API.LogError(">>> [에러] failed to ensure bot account: " + err.Error())
		return ""
	}

	return botUserID
}

func (p *Plugin) settingScheduler() *cluster.Job {
	job, err := cluster.Schedule(
		p.API,
		"BackgroundJob",
		cluster.MakeWaitForRoundedInterval(24*time.Hour),
		func() {
			p.API.LogInfo("Start Scheduler Job")
			jobErr := SendNotification(p)

			if jobErr != nil {
				p.API.LogError(">>> [에러] Failed to send notification. Error: " + jobErr.Error())
			}
		},
	)

	if err != nil {
		p.API.LogError(">>> [에러] Unable to schedule job for standup reports: " + err.Error())
		return nil
	}
	return job
}
