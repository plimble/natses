package natses

import (
	"github.com/golang/protobuf/proto"
	"github.com/plimble/natses/pb"
)

type Marshaler interface {
	Marshal() ([]byte, error)
}

func NewEvent(eventType string, data proto.Message) ([]byte, error) {
	bdata, err := proto.Marshal(data)
	if err != nil {
		return nil, err
	}

	epb := &pb.Event{
		Type: eventType,
		Data: bdata,
	}

	return epb.Marshal()
}
