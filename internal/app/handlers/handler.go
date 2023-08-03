package handlers

import (
	"net/http"

	"api-gingonic/internal/app/models"
	"api-gingonic/internal/app/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type HandlerDependency struct {
	service service.Transaction
}

func NewHandlerDependency(service service.Transaction) *HandlerDependency {
	return &HandlerDependency{service: service}
}

type TransactionService interface {
	CreateTransaction(transaction models.Transaction) error
}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *HandlerDependency) TransactionHandler(c *gin.Context) {
	transaction := models.Transaction{Operation: "Purchase", Amount: 100.00, PaymentId: "1234"}

	err := h.service.CreateTransaction(transaction)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"object": transaction,
	})
}

func CreateTransactionHandler(c *gin.Context) {
	var transaction models.Transaction

	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		panic(err)
	}

	if err := validateStruct(transaction); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"CreatedTrasaction": transaction,
	})

}

func validateStruct(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		return err
	}
	return nil
}
