package natsclient

import (
	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	conn *nats.Conn
}

func NewSubscriber(url string) (*Subscriber, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &Subscriber{conn: nc}, nil
}

func (s *Subscriber) Subscribe(subject string, handler func(msg []byte)) error {
	_, err := s.conn.Subscribe(subject, func(m *nats.Msg) {
		handler(m.Data)
	})
	if err != nil {
		return err
	}
	return nil
}
