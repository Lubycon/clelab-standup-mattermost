package main

import (
	"time"

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

	return nil
}
