package fmt

import (
	"fmt"
	"github.com/KirillBogatikov/Cuba/log/go/common"
)

type Stream struct {
	Formatter Formatter
}

func NewStream(formatter *Formatter) Stream {
	stream := Stream{}
	if formatter == nil {
		stream.Formatter = DefaultFormatter{}
	} else {
		stream.Formatter = *formatter
	}

	return stream
}

func (t Stream) Write(record common.Record) {
	text := t.Formatter.Format(record)
	fmt.Println(text)
}
