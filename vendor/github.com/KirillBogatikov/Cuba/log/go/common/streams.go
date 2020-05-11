package common

type LogStream interface {
	Write(record Record)
}
