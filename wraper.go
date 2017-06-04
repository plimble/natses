package natses

import (
	"github.com/golang/protobuf/proto"
	nats "github.com/nats-io/go-nats"
	"github.com/nats-io/go-nats-streaming"
)

//go:generate mockery -name Conn
type Conn interface {
	// Publish
	Publish(subject, eventType string, data proto.Message) error
	PublishAsync(subject, eventType string, data proto.Message, ah stan.AckHandler) (string, error)

	// Subscribe
	Subscribe(subject string, cb MsgHandler, opts ...stan.SubscriptionOption) (stan.Subscription, error)

	// QueueSubscribe
	QueueSubscribe(subject, qgroup string, cb MsgHandler, opts ...stan.SubscriptionOption) (stan.Subscription, error)

	// Close
	Close() error

	// NatsConn returns the underlying NATS conn. Use this with care. For
	// example, closing the wrapped NATS conn will put the NATS Streaming Conn
	// in an invalid state.
	NatsConn() *nats.Conn
}

type ESNats struct {
	conn stan.Conn
}

func New(conn stan.Conn) Conn {
	return &ESNats{conn}
}

func (n *ESNats) Publish(subject, eventType string, data proto.Message) error {
	edata, err := NewEvent(eventType, data)
	if err != nil {
		return err
	}

	return n.conn.Publish(subject, edata)
}

func (n *ESNats) PublishAsync(subject, eventType string, data proto.Message, ah stan.AckHandler) (string, error) {
	edata, err := NewEvent(eventType, data)
	if err != nil {
		return "", err
	}

	return n.conn.PublishAsync(subject, edata, ah)
}

func (n *ESNats) Subscribe(subject string, cb MsgHandler, opts ...stan.SubscriptionOption) (stan.Subscription, error) {
	return n.conn.Subscribe(subject, SubEvent(cb), opts...)
}

func (n *ESNats) QueueSubscribe(subject, qgroup string, cb MsgHandler, opts ...stan.SubscriptionOption) (stan.Subscription, error) {
	return n.conn.QueueSubscribe(subject, qgroup, SubEvent(cb), opts...)
}

func (n *ESNats) Close() error {
	return n.conn.Close()
}

func (n *ESNats) NatsConn() *nats.Conn {
	return n.conn.NatsConn()
}
