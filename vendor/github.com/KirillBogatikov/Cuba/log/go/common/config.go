package common

type Configuration struct {
	Debug LogStream
	Info  LogStream
	Warn  LogStream
	Error LogStream
}
