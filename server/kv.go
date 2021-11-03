package main

import (
	"encoding/json"

	"github.com/mattermost/mattermost-plugin-starter-template/server/types"
)

func getChannels(p *Plugin) types.ChannelList {
	channelListData, _ := p.API.KVGet(ChannelListKey)
	channelList := types.ChannelList{}
	err := json.Unmarshal(channelListData, &channelList)
	if err != nil {
		p.API.LogError(">>> [에러] unmarshal error: " + err.Error())
	}
	return channelList
}
