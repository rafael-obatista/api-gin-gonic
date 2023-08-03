package repository

import (
	"api-gingonic/internal/app/models"

	"gorm.io/gorm"
)

type PostgresMock struct{}

func (p *PostgresMock) Connection() error {
	return nil
}

func (p *PostgresMock) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	return transaction, nil
}

//func TestPostgres_CreateTransaction(t *testing.T) {
//
//	post := PostgresMock{}
//	a := NewPostgres(&post)
//
//
//
//}

type MockDatabase struct {
	// Implemente as funções da interface Database de acordo com o comportamento desejado no teste.
}

func (m MockDatabase) Create(i interface{}) *gorm.DB {
	//TODO implement me

	return nil
}

//func TestCreateTransaction(t *testing.T) {
//	mockDB := &MockDatabase{}
//	//repo := NewPostgres(mockDB)
//
//	transaction := models.Transaction{Operation: "Purchase", Amount: 100.00, PaymentId: "1234"}
//
//	createTransaction, err := repo.CreateTransaction(transaction)
//	if err != nil {
//		return
//	}
//
//	assert.Nil(t, createTransaction)
//
//	// Realize os testes utilizando o repo.
//}

//	a.Connection()
//
//
//
//
//	transaction := models.Transaction{{},
//
//	//func NewPostgres(options ...*Postgres) Transaction {
//	//	a := &Postgres{}
//	//	if len(options) > 0 {
//	//	a = options[0]
//	//}
//	//
//	//	a.Connection()
//	//	return a
//	//}
//
//}
