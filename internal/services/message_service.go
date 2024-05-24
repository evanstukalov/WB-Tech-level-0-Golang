package services

import (
	"log"

	"github.com/evanstukalov/wildberries_internship_l0/internal/cache"
	"github.com/evanstukalov/wildberries_internship_l0/internal/database"
	"github.com/evanstukalov/wildberries_internship_l0/internal/validation"
)

type MessageService struct {
	cache    *cache.Cache
	database *database.Database
}

func NewMessageService(cache *cache.Cache, db *database.Database) *MessageService {
	return &MessageService{cache: cache, database: db}
}

func (messageService MessageService) ProcessMessage(message []byte) error {

	order, err := validation.ValidateOrderJSON(string(message))

	if err != nil {
		log.Printf("Validation error: %v", err)
		return err
	}

	err = (*messageService.cache).Add(order.OrderUID, order)
	if err != nil {
		log.Printf("Cache error: %v", err)
		return err
	}

	// db

	return nil

}
