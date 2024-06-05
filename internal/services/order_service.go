package services

import (
	"github.com/evanstukalov/wildberries_internship_l0/internal/models"
	"log"

	"github.com/evanstukalov/wildberries_internship_l0/internal/cache"
	"github.com/evanstukalov/wildberries_internship_l0/internal/database"
	"github.com/evanstukalov/wildberries_internship_l0/internal/utils"
)

type OrderService interface {
	ProcessOrder(data []byte) error
}

type OrderServiceImpl struct {
	cache    cache.Cache
	database database.ServiceDataBase
}

func NewMessageService(cache cache.Cache, db database.ServiceDataBase) *OrderServiceImpl {
	return &OrderServiceImpl{cache: cache, database: db}
}

func (orderService OrderServiceImpl) ProcessOrder(message []byte) error {

	order, err := utils.ValidateOrderJSON(string(message))

	if err != nil {
		log.Printf("Validation error: %v", err)
		return err
	}

	err = orderService.database.Create(order)
	if err != nil {
		log.Printf("Database error: %v", err)
		return err
	}

	err = orderService.cache.Set(order.OrderUID, order)
	if err != nil {
		log.Printf("Cache error: %v", err)
		return err
	}

	return nil

}

func (orderService OrderServiceImpl) FillCacheWithOrders() error {
	var orders map[string]models.Order

	orders, err := orderService.database.GetAll()
	if err != nil {
		return err
	}

	ordersInterface := make(map[string]interface{})
	for key, value := range orders {
		ordersInterface[key] = value
	}

	err = orderService.cache.FillUp(ordersInterface)
	if err != nil {
		return err
	}

	return nil
}
