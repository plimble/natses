package natses

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/plimble/natses/pb"
)

func NewEvent(eventType string, data proto.Message) ([]byte, error) {
	bdata, err := proto.Marshal(data)
	if err != nil {
		return nil, err
	}

	epb := &pb.Event{
		Type: eventType,
		Data: bdata,
		Time: time.Now().Unix(),
	}

	return proto.Marshal(epb)
}
