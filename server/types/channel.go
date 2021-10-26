package types

type ChannelList = []Channel

type Channel struct {
	ID    string   `json:"id"`
	Users []string `json:"users"`
}
