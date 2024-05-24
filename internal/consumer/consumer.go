package consumer

import (
	"github.com/nats-io/nats.go"
)

type Consumer struct {
	natsConn *nats.Conn
	// processor *Processor
}

func NewSubscriber(natsURL string) (*Consumer, error) {

	nc, err := nats.Connect(natsURL)

	if err != nil {
		return nil, err
	}

	return &Consumer{
		natsConn: nc,
	}, nil
}

func (s *Consumer) Close() {
	if s.natsConn != nil {
		s.natsConn.Close()
	}
}

func (s *Consumer) Subscribe(subject string) error {
	_, err := s.natsConn.Subscribe(subject, func(msg *nats.Msg) {

		// Validation

		// ProcessData

	})

	return err
}
