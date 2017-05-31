package natses

import (
	"github.com/nats-io/go-nats-streaming"
	"github.com/plimble/natses/pb"
)

type EventMsg struct {
	*stan.Msg
	EventType string
	Data      []byte
}

type MsgHandler func(msg *EventMsg)

func SubEvent(h MsgHandler) stan.MsgHandler {
	return func(msg *stan.Msg) {
		pbe := &pb.Event{}
		pbe.Unmarshal(msg.Data)

		h(&EventMsg{
			Msg:       msg,
			EventType: pbe.Type,
			Data:      pbe.Data,
		})
	}
}
