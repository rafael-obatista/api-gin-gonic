package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Operation string  `json:"operation" validate:"required"`
	Amount    float64 `json:"amount" validate:"gte=0"`
	PaymentId string  `json:"payment_id"`
}
