package natses

import (
	"github.com/plimble/natses/pb"
)

type Marshaler interface {
	Marshal() ([]byte, error)
}

func NewEvent(eventType string, data Marshaler) ([]byte, error) {
	bdata, err := data.Marshal()
	if err != nil {
		return nil, err
	}

	epb := &pb.Event{
		Type: eventType,
		Data: bdata,
	}

	bepb, err := epb.Marshal()
	if err != nil {
		return nil, err
	}

	return bepb, nil
}
