package models

import (
	"time"
)

type Transaction struct {
	TransactionID   uint      `json:"transaction_id" gorm:"primaryKey"`
	AccountID       uint      `json:"account_id"`
	OperationTypeID uint      `json:"operation_type_id"`
	Amount          int       `json:"amount"` // Stored as integer (cents)
	EventDate       time.Time `json:"event_date" gorm:"autoCreateTime"`
}
