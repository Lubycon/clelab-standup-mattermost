package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mattermost/mattermost-plugin-api/cluster"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

type Plugin struct {
	plugin.MattermostPlugin

	router *mux.Router

	emptyTime time.Time

	ServerConfig *model.Config

	userID string

	job *cluster.Job
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
