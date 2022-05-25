package models

type NewGame struct {
	Author              string   `json:"author"`
	ChannelID           string   `json:"channel_id"`
	VoiceChannelMembers []string `json:"voice_channel_members"`
}
