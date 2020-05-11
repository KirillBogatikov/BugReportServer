package fmt

import (
	"encoding/json"
	"fmt"
	"github.com/KirillBogatikov/Cuba/log/go/common"
)

type Formatter interface {
	Format(record common.Record) string
}

type DefaultFormatter struct {
	UseJson bool
}

func (t DefaultFormatter) Format(record common.Record) string {
	payload := ""

	if record.Data != nil {
		if t.UseJson {
			bytes, err := json.Marshal(record.Data)
			if err == nil {
				payload = string(bytes)
			} else {
				payload = fmt.Sprintf("\n{\"error\": \"%s\", \"fmt\": \"%s\"}", err, record.Data)
			}
		} else {
			payload = fmt.Sprintf("\n%s", record.Data)
		}
	}

	if record.Error != nil {
		if len(payload) > 0 {
			payload += "\n"
		}

		payload += fmt.Sprint(record.Error)
	}

	return fmt.Sprintf("[%s] %s [%s]: %s %s",
		record.Level.Name, record.Time.Format("2006-01-02 15:04:05.000"), record.Tag, record.Message, payload)
}
