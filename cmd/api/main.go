package main

import (
	"api-gingonic/internal/app/handlers"
	"api-gingonic/internal/app/repository"
	"api-gingonic/internal/app/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := repository.NewPostgres()
	transactionService := service.NewTransactionService(repo)
	a := handlers.NewHandlerDependency(transactionService)

	router.GET("/ping", handlers.PingHandler)
	router.GET("/transaction", a.TransactionHandler)
	router.POST("/transaction", handlers.CreateTransactionHandler)

	router.Run(":8082")
}
