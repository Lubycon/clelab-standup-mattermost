package main

import (
	"time"

	"github.com/pkg/errors"
)

func (p *Plugin) OnActivate() error {
	p.ServerConfig = p.API.GetConfig()
	p.router = p.InitAPI()
	p.emptyTime = time.Time{}.AddDate(1, 1, 1)

	if p.registerCommand() != nil {
		return errors.Wrap(p.registerCommand(), "failed to register command")
	}

	return nil
}
