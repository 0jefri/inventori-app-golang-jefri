package dto

import (
	"time"

	"github.com/inventori-app-jeff/internal/model"
)

type TransactionResponse struct {
	ID              string        `json:"id"`
	Product         model.Product `json:"product"`
	TransactionType string        `json:"transactionType"`
	Amount          float64       `json:"amount"`
	Description     string        `json:"description"`
	Timestamp       time.Time     `json:"timestamp"`
}
