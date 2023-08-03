package service

import (
	"fmt"

	"api-gingonic/internal/app/models"
	"api-gingonic/internal/app/repository"
)

type Transaction interface {
	CreateTransaction(transaction models.Transaction) error
	DeleteTransaction(transaction models.Transaction) error
}

func NewTransactionService(repo repository.Transaction) Transaction {
	return &transaction{repository: repo}
}

type transaction struct {
	repository repository.Transaction
}

func (t *transaction) CreateTransaction(transaction models.Transaction) error {
	fmt.Println("Passando pelo service", transaction)

	_, err := t.repository.CreateTransaction(transaction)
	if err != nil {
		return err
	}
	fmt.Println("Retorno do repository", err)
	return nil
}

func (t *transaction) DeleteTransaction(transaction models.Transaction) error {
	fmt.Println("real service transacion", transaction)
	fmt.Println("chamar o repository aqui dentro")
	return nil
}
