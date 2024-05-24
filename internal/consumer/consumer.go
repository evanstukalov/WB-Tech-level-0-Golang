package consumer

import (
	"log"

	"github.com/evanstukalov/wildberries_internship_l0/internal/services"
	"github.com/nats-io/nats.go"
)

type Consumer struct {
	natsConn *nats.Conn
	service  *services.MessageService
}

func NewConsumer(natsURL string, service *services.MessageService) (*Consumer, error) {

	nc, err := nats.Connect(natsURL)

	if err != nil {
		return nil, err
	}

	return &Consumer{
		natsConn: nc,
		service:  service,
	}, nil
}

func (s *Consumer) Close() {
	if s.natsConn != nil {
		s.natsConn.Close()
	}
}

func (s *Consumer) Consume(subject string) error {
	_, err := s.natsConn.Subscribe(subject, func(msg *nats.Msg) {

		if err := s.service.ProcessMessage(msg.Data); err != nil {
			log.Printf("Error processing message: %v", err)
		}

	})

	return err
}
