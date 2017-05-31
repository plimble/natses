package natses

import (
	"github.com/plimble/natses/pb"
)

type Marshaler interface {
	Marshal() ([]byte, error)
}

func NewEvent(eventType string, data Marshaler) []byte {
	bdata, _ := data.Marshal()
	epb := &pb.Event{
		Type: eventType,
		Data: bdata,
	}

	bepb, _ := epb.Marshal()
	return bepb
}
