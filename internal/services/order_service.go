package services

import (
	"github.com/evanstukalov/wildberries_internship_l0/internal/models"
	"log"

	"github.com/evanstukalov/wildberries_internship_l0/internal/cache"
	"github.com/evanstukalov/wildberries_internship_l0/internal/database"
	"github.com/evanstukalov/wildberries_internship_l0/internal/validation"
)

type OrderService struct {
	cache    *cache.Cache
	database *database.Database
}

func NewMessageService(cache *cache.Cache, db *database.Database) *OrderService {
	return &OrderService{cache: cache, database: db}
}

func (orderService OrderService) ProcessOrder(message []byte) error {

	order, err := validation.ValidateOrderJSON(string(message))

	if err != nil {
		log.Printf("Validation error: %v", err)
		return err
	}

	err = orderService.database.CreateOrder(order)
	if err != nil {
		log.Printf("Database error: %v", err)
		return err
	}

	err = (*orderService.cache).Add(order.OrderUID, order)
	if err != nil {
		log.Printf("Cache error: %v", err)
		return err
	}

	return nil

}

func (orderService OrderService) FillCacheWithOrders() error {
	var orders map[string]models.Order

	orders, err := (*orderService.database).GetOrders()
	if err != nil {
		return err
	}

	err = (*orderService.cache).FillUp(orders)
	if err != nil {
		return err
	}

	return nil
}
