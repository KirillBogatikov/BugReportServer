package common

import (
	"time"
)

type Level struct {
	Name     string
	Priority int
}

var (
	DEBUG = Level{Name: "DEBUG", Priority: 0}
	INFO  = Level{Name: "INFO", Priority: 1}
	WARN  = Level{Name: "WARN", Priority: 2}
	ERROR = Level{Name: "ERROR", Priority: 3}
)

type Record struct {
	Level   Level
	Tag     string
	Message string
	Time    time.Time
	Error   error
	Data    interface{}
}
