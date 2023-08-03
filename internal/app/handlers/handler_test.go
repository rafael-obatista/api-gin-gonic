package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"api-gingonic/internal/app/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MyTransactionService struct {
}

func (s *MyTransactionService) DeleteTransaction(transaction models.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (s *MyTransactionService) CreateTransaction(transaction models.Transaction) error {
	fmt.Println("teste mock", transaction)
	return fmt.Errorf("error")
}

func TestTransactionHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	new := NewHandlerDependency(&MyTransactionService{})

	router.GET("/transaction", new.TransactionHandler)

	req, err := http.NewRequest("GET", "/transaction", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, `{"object":{"operation":"Purchase","amount":100,"payment_id":"1234"}}`, recorder.Body.String())

}
