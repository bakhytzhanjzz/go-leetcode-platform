package natsclient

import (
	"log"

	"github.com/nats-io/nats.go"
)

type Publisher struct {
	conn *nats.Conn
}

func NewPublisher(url string) (*Publisher, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &Publisher{conn: nc}, nil
}

func (p *Publisher) Publish(subject string, msg []byte) {
	if err := p.conn.Publish(subject, msg); err != nil {
		log.Printf("Failed to publish message to subject %s: %v", subject, err)
	}
}
