package reference

type User struct {
	Id        string `json:"id"`
	ChannelID string `json:"channel_id"`
	Afk       bool   `json:"afk"`
}

type UserChannel struct {
}
