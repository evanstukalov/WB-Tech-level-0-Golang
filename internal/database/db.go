package database

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/evanstukalov/wildberries_internship_l0/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() (*Database, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return nil, err
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Println("DATABASE_URL is not set")
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}

	log.Println("Database connection established successfully")
	return &Database{db: db}, nil
}

func (d *Database) CreateOrder(order *models.Order) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order.Delivery).Error; err != nil {
			return err
		}
		order.DeliveryID = order.Delivery.DeliveryID

		if err := tx.Create(&order.Payment).Error; err != nil {
			return err
		}
		order.PaymentID = order.Payment.PaymentID

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		for i := range order.Items {
			order.Items[i].OrderUID = order.OrderUID
			if err := tx.Create(&order.Items[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (d *Database) GetOrders() (map[string]models.Order, error) {
	var orders []models.Order

	if err := d.db.Find(&orders).Error; err != nil {
		log.Println("Error retrieving orders:", err)
		return nil, err
	}

	for i := range orders {

		var delivery models.Delivery
		var payment models.Payment
		var item models.Item

		d.db.Take(&delivery, "delivery_id = ?", orders[i].DeliveryID)
		d.db.Take(&payment, "payment_id = ?", orders[i].PaymentID)
		d.db.Take(&item, "order_uid = ?", orders[i].OrderUID)

		orders[i].Delivery = delivery
		orders[i].Payment = payment
		orders[i].Items = append(orders[i].Items, item)

	}

	orderMap := make(map[string]models.Order)

	for _, order := range orders {
		orderMap[order.OrderUID] = order
	}

	return orderMap, nil
}
