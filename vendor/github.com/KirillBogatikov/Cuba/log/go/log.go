package go_log

import (
	"github.com/KirillBogatikov/Cuba/log/go/common"
	"time"
)

type Log struct {
	Configuration common.Configuration
}

func NewLog(configuration common.Configuration) Log {
	return Log{
		Configuration: configuration,
	}
}

func (t Log) Write(level common.Level, tag, msg string, err error, data interface{}) {
	var stream common.LogStream
	switch level {
	case common.DEBUG:
		stream = t.Configuration.Debug
		break
	case common.INFO:
		stream = t.Configuration.Info
		break
	case common.WARN:
		stream = t.Configuration.Warn
		break
	case common.ERROR:
		stream = t.Configuration.Error
		break
	default:
		return
	}

	if len(tag) == 0 {
		tag = "GO"
	}

	stream.Write(common.Record{
		Level:   level,
		Tag:     tag,
		Message: msg,
		Time:    time.Now(),
		Error:   err,
		Data:    data,
	})
}

func (t Log) Debug1(message string) {
	t.Write(common.DEBUG, "", message, nil, nil)
}

func (t Log) Debug2(tag, message string) {
	t.Write(common.DEBUG, tag, message, nil, nil)
}

func (t Log) Debug3(tag, message string, data interface{}) {
	t.Write(common.DEBUG, tag, message, nil, data)
}

func (t Log) DebugE(tag, message string, err error) {
	t.Write(common.DEBUG, tag, message, err, nil)
}

func (t Log) Info1(message string) {
	t.Write(common.INFO, "", message, nil, nil)
}

func (t Log) Info2(tag, message string) {
	t.Write(common.INFO, tag, message, nil, nil)
}

func (t Log) Info3(tag, message string, data interface{}) {
	t.Write(common.INFO, tag, message, nil, data)
}

func (t Log) InfoE(tag, message string, err error) {
	t.Write(common.INFO, tag, message, err, nil)
}

func (t Log) Warm1(message string) {
	t.Write(common.WARN, "", message, nil, nil)
}

func (t Log) Warn2(tag, message string) {
	t.Write(common.WARN, tag, message, nil, nil)
}

func (t Log) Warn3(tag, message string, data interface{}) {
	t.Write(common.WARN, tag, message, nil, data)
}

func (t Log) WarnE(tag, message string, err error) {
	t.Write(common.WARN, tag, message, err, nil)
}

func (t Log) Error1(message string) {
	t.Write(common.ERROR, "", message, nil, nil)
}

func (t Log) Error2(tag, message string) {
	t.Write(common.ERROR, tag, message, nil, nil)
}

func (t Log) Error3(tag, message string, data interface{}) {
	t.Write(common.ERROR, tag, message, nil, data)
}

func (t Log) ErrorE(tag, message string, err error) {
	t.Write(common.ERROR, tag, message, err, nil)
}
