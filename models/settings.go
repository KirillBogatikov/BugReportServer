package models

type BotSettings struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	Password string `json:"password"`
}

type Settings struct {
	Bots    []BotSettings `json:"bots"`
	Clients []string      `json:"clients"`
}
