package database

import (
	"fmt"

	"github.com/evanstukalov/wildberries_internship_l0/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDataBase(dataSourceName string) (*Database, error) {

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil, err
	}

	/*	err = db.AutoMigrate(&models.Order{}, &models.Delivery{}, &models.Payment{}, &models.Item{})
		if err != nil {
			fmt.Println("Error creating database table:", err)
		}*/

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
