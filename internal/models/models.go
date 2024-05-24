package models

type Order struct {
	OrderUID          string   `validate:"required"`
	TrackNumber       string   `validate:"required"`
	Entry             string   `validate:"required"`
	Delivery          Delivery `validate:"required"`
	Payment           Payment  `validate:"required"`
	Items             []Item   `validate:"required,dive"`
	Locale            string   `validate:"required"`
	InternalSignature string   `validate:"required"`
	CustomerID        string   `validate:"required"`
	DeliveryService   string   `validate:"required"`
	ShardKey          string   `validate:"required"`
	SmID              int      `validate:"required"`
	DateCreated       string   `validate:"required"`
	OofShard          string   `validate:"required"`
}

type Delivery struct {
	Name    string `validate:"required"`
	Phone   string `validate:"required"`
	Zip     string `validate:"required"`
	City    string `validate:"required"`
	Address string `validate:"required"`
	Region  string `validate:"required"`
	Email   string `validate:"required,email"`
}

type Payment struct {
	Transaction  string  `validate:"required"`
	RequestID    string  `validate:"required"`
	Currency     string  `validate:"required"`
	Provider     string  `validate:"required"`
	Amount       float64 `validate:"required,gte=0"`
	PaymentDT    int64   `validate:"required,gte=0"`
	Bank         string  `validate:"required"`
	DeliveryCost int     `validate:"required,gte=0"`
	GoodsTotal   int     `validate:"required,gte=0"`
	CustomFee    int     `validate:"required,gte=0"`
}

type Item struct {
	ChrtID      int    `validate:"required"`
	TrackNumber string `validate:"required"`
	Price       int    `validate:"required,gte=0"`
	Rid         string `validate:"required"`
	Name        string `validate:"required"`
	Sale        int    `validate:"required,gte=0"`
	Size        string `validate:"required"`
	TotalPrice  int    `validate:"required,gte=0"`
	NmID        int    `validate:"required"`
	Brand       string `validate:"required"`
	Status      int    `validate:"required,gte=0"`
}
