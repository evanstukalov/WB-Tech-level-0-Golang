package models

type Order struct {
	OrderUID          string   `json:"order_uid" validate:"required" gorm:"primaryKey"`
	TrackNumber       string   `json:"track_number"`
	Entry             string   `json:"entry"`
	Delivery          Delivery `json:"delivery" gorm:"-"`
	Payment           Payment  `json:"payment" gorm:"-"`
	Items             []Item   `json:"items" validate:"required,dive" gorm:"-"`
	Locale            string   `json:"locale"`
	InternalSignature string   `json:"internal_signature"`
	CustomerID        string   `json:"customer_id"`
	DeliveryService   string   `json:"delivery_service"`
	ShardKey          string   `json:"shardkey"`
	SmID              int      `json:"sm_id"`
	DateCreated       string   `json:"date_created"`
	OofShard          string   `json:"oof_shard"`
	DeliveryID        int      `json:"delivery_id"`
	PaymentID         string   `json:"payment_id"`
}

type Delivery struct {
	DeliveryID int    `gorm:"primaryKey;autoIncrement"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Zip        string `json:"zip"`
	City       string `json:"city"`
	Address    string `json:"address"`
	Region     string `json:"region"`
	Email      string `json:"email" validate:"email"`
}

type Payment struct {
	PaymentID    string  `gorm:"primaryKey;autoIncrement"`
	Transaction  string  `json:"transaction"`
	RequestID    string  `json:"request_id"`
	Currency     string  `json:"currency"`
	Provider     string  `json:"provider"`
	Amount       float64 `json:"amount" validate:"gte=0"`
	PaymentDT    int64   `json:"payment_dt" validate:"gte=0"`
	Bank         string  `json:"bank"`
	DeliveryCost int     `json:"delivery_cost" validate:"gte=0"`
	GoodsTotal   int     `json:"goods_total" validate:"gte=0"`
	CustomFee    int     `json:"custom_fee" validate:"gte=0"`
}

type Item struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price" validate:"gte=0"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale" validate:"gte=0"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price" validate:"gte=0"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status" validate:"gte=0"`
	OrderUID    string `gorm:"primaryKey"`
}
