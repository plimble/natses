package natses

import (
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats-streaming"
	"github.com/plimble/natses/pb"
)

type EventMsg struct {
	*stan.Msg
	EventType string
	Data      []byte
	Timestamp int64
}

type MsgHandler func(msg *EventMsg)

func SubEvent(h MsgHandler) stan.MsgHandler {
	return func(msg *stan.Msg) {
		pbe := &pb.Event{}
		proto.Unmarshal(msg.Data, pbe)

		h(&EventMsg{
			Msg:       msg,
			EventType: pbe.Type,
			Data:      pbe.Data,
			Timestamp: msg.Timestamp,
		})
	}
}
