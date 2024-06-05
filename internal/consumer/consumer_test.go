package consumer

import (
	"github.com/nats-io/nats.go"
	"testing"
)

type MockOrderService struct {
	ProcessOrderFn func(data []byte) error
}

func (m *MockOrderService) ProcessOrder(data []byte) error {
	return m.ProcessOrderFn(data)
}

func TestNewConsumer(t *testing.T) {
	mockService := &MockOrderService{}
	natsURL := nats.DefaultURL

	consumer, err := NewConsumer(natsURL, mockService)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if consumer.natsConn == nil {
		t.Fatal("Expected natsConn to be initialized")
	}

	if consumer.service != mockService {
		t.Fatal("Expected service to be set correctly")
	}
}

func TestConsume(t *testing.T) {
	mockService := &MockOrderService{
		ProcessOrderFn: func(data []byte) error {
			if string(data) != "test message" {
				t.Fatalf("Expected 'test message', got %s", data)
			}
			return nil
		},
	}
	natsURL := nats.DefaultURL

	consumer, err := NewConsumer(natsURL, mockService)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	subject := "test.subject"

	err = consumer.Consume(subject)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	nc := consumer.natsConn
	err = nc.Publish(subject, []byte("test message"))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	nc.Flush()

	if err := nc.LastError(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
