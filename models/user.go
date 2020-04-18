package models

import "encoding/json"

type Reviewer struct {
	Id   int
	Name string
	Info interface{}
}

type TelegramInfo struct {
	ChatId int64
}

type InfoMapper func(bytes []byte) interface{}

func TelegramInfoMapper(bytes []byte) interface{} {
	telegramInfo := &TelegramInfo{}
	json.Unmarshal(bytes, telegramInfo)
	return *telegramInfo
}

func NewTelegramReviewer(name string, chatId int64) *Reviewer {
	return &Reviewer{
		Name: name,
		Info: TelegramInfo{chatId},
	}
}
