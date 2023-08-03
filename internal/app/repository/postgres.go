package repository

import (
	"fmt"

	"api-gingonic/internal/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//func WithDatabase(database *gorm.DB) Options {
//	return func(o *Postgres) {
//		o.database = database
//	}
//}
//
//type Options func(*Postgres)

func NewPostgres() Transaction {
	a := &Postgres{}
	a.Connection()
	return a
}

type Postgres struct {
	database *gorm.DB
}

func (p *Postgres) Connection() error {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error", err)
		panic("failed to connect database")
	}

	p.database = db

	err = db.AutoMigrate(&models.Transaction{})
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	fmt.Println("Chegou no repository", transaction)

	if p.database == nil {
		return transaction, fmt.Errorf("database is null")
	}

	result := p.database.Create(&transaction)
	if result.Error != nil {
		return transaction, fmt.Errorf("error creating transaction: %w", result.Error)
	}

	fmt.Println("transaction created successfully", result.RowsAffected)

	return transaction, nil
}
