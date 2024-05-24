package validation

import (
	"encoding/json"
	"errors"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateOrderJSON(jsonStr string) (*Order, error) {
	var order Order

	err := json.Unmarshal([]byte(jsonStr), &order)
	if err != nil {
		return nil, errors.New("error parsing JSON: " + err.Error())
	}

	err = validate.Struct(order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
