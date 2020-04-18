package models

import "time"

type Reporter struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}

type Error struct {
	Line       int    `json:"line"`
	Method     string `json:"method"`
	Class      string `json:"class"`
	StackTrace string `json:"trace"`
}

type Report struct {
	Author     Reporter  `json:"user"`
	Message    string    `json:"message"`
	Error      *Error    `json:"error"`
	SendTime   time.Time `json:"time"`
	ServedTime time.Time
	AppName    string
}
