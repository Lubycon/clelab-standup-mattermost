package main

import (
	"encoding/json"

	"github.com/mattermost/mattermost-plugin-starter-template/server/types"
)

func getChannels(p *Plugin) types.ChannelList {
	kv, _ := p.API.KVGet(ChannelListKey)
	channels := types.ChannelList{}
	err := json.Unmarshal(kv, &channels)
	if err != nil {
		p.API.LogError(">>> [에러] unmarshal error: " + err.Error())
	}

	return channels
}
