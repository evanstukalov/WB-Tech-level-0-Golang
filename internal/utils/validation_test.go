package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateOrderJSON_ValidJSON(t *testing.T) {
	validate = validator.New()

	jsonStr := `{
        "order_uid": "12345",
        "track_number": "TN12345",
        "entry": "entry1",
        "delivery": {
            "name": "John Doe",
            "phone": "+1234567890",
            "zip": "12345",
            "city": "City",
            "address": "Address",
            "region": "Region",
            "email": "email@example.com"
        },
        "payment": {
            "transaction": "txn12345",
            "request_id": "",
            "currency": "USD",
            "provider": "provider1",
            "amount": 100,
            "payment_dt": 1633024800,
            "bank": "bank1",
            "delivery_cost": 10,
            "goods_total": 90,
            "custom_fee": 0
        },
        "items": [
            {
                "chrt_id": 123,
                "track_number": "TN12345",
                "price": 100,
                "rid": "rid1",
                "name": "item1",
                "sale": 10,
                "size": "M",
                "total_price": 90,
                "nm_id": 456,
                "brand": "brand1",
                "status": 1
            }
        ],
        "locale": "en",
        "internal_signature": "",
        "customer_id": "cust12345",
        "delivery_service": "service1",
        "shardkey": "",
        "sm_id": 1,
        "date_created": "2021-09-30T15:00:00Z",
        "oof_shard": "",
        "delivery_id": 1,
        "payment_id": ""
    }`

	order, err := ValidateOrderJSON(jsonStr)
	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, order.OrderUID, "12345")
}

func TestValidateOrderJSON_MissingRequiredFields(t *testing.T) {
	validate = validator.New()

	jsonStr := `{
        "track_number": "TN12345",
        "entry": "",
        "delivery": {
            "name": "",
            "phone": "",
            "zip": "",
            "city": "",
            "address": "",
            "region": "",
            "email": ""
        },
        "payment": {
            "transaction": "",
            "request_id": "",
            "currency": "",
            "provider": "",
            "amount": 0,
            "payment_dt": 0,
            "bank": "",
            "delivery_cost": 0,
            "goods_total": 0,
            "custom_fee": 0
        },
        "items": [],
        "locale": "",
        "internal_signature": "",
        "customer_id": "",
        "delivery_service": "",
        "shardkey": "",
        "sm_id": 0,
        "date_created": "",
        "oof_shard": "",
        "delivery_id": 0,
        "payment_id": ""
    }`

	order, err := ValidateOrderJSON(jsonStr)
	assert.Error(t, err)
	assert.Nil(t, order)
}
