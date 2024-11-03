package model

import "time"

type Transaction struct {
	ID              string    `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	ProductID       string    `gorm:"type:uuid;not null;references:ID" json:"productID" binding:"required"`
	TransactionType string    `gorm:"varchar(255);not null" json:"transactionType" binding:"required,alpha,oneof=sendProduct receiveProduct"`
	Amount          float64   `gorm:"type:float;not null;default:0;check:Amount >= 0" json:"amount" binding:"required,numeric"`
	Description     string    `gorm:"type:text" json:"description" binding:"omitempty"`
	Timestamp       time.Time `gorm:"type:timestamp;default:current_timestamp" json:"timestamp"`
}
