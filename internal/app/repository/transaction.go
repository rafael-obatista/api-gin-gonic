package repository

import "api-gingonic/internal/app/models"

type Transaction interface {
	Connection() error
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
}
