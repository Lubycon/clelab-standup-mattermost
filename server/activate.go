package main

import (
	"time"

	"github.com/mattermost/mattermost-plugin-api/cluster"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/pkg/errors"
)

func (p *Plugin) OnActivate() error {
	p.ServerConfig = p.API.GetConfig()
	p.router = p.InitAPI()
	p.emptyTime = time.Time{}.AddDate(1, 1, 1)

	botUserID, err := p.Helpers.EnsureBot(&model.Bot{

		Username:    "cletandup",
		DisplayName: "Cletandup",
		Description: "Hello! I'm standup bot for clelab.",
	})

	if err != nil {
		return errors.Wrap(err, "failed to ensure bot account")
	}
	p.userID = botUserID

	if p.registerCommand() != nil {
		return errors.Wrap(p.registerCommand(), "failed to register command")
	}

	job, err := cluster.Schedule(
		p.API,
		"BackgroundJob",
		cluster.MakeWaitForRoundedInterval(24*time.Hour),
		func() {
			p.API.LogInfo("Start Scheduler Job")
			jobErr := SendNotification(p)

			if jobErr != nil {
				p.API.LogError(">>> [에러] Failed to send notification. Error: " + err.Error())
			}
		},
	)

	if err != nil {
		p.API.LogError(">>> [에러] Unable to schedule job for standup reports: " + err.Error())
		return err
	}

	p.job = job

	return nil
}
