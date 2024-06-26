package consumer

import (
	"log"

	"github.com/evanstukalov/wildberries_internship_l0/internal/services"
	"github.com/nats-io/nats.go"
)

type Consumer struct {
	natsConn *nats.Conn
	service  services.OrderService
}

func NewConsumer(natsURL string, service services.OrderService) (*Consumer, error) {

	nc, err := nats.Connect(natsURL)

	if err != nil {
		return nil, err
	}

	return &Consumer{
		natsConn: nc,
		service:  service,
	}, nil
}

func (s *Consumer) Consume(subject string) error {
	_, err := s.natsConn.Subscribe(subject, func(msg *nats.Msg) {

		if err := s.service.ProcessOrder(msg.Data); err != nil {
			log.Printf("Error processing message: %v", err)
		}

	})

	return err
}
