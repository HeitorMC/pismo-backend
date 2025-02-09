package models

type Account struct {
	AccountID      uint          `json:"account_id" gorm:"primaryKey"`
	DocumentNumber string        `json:"document_number" gorm:"unique;not null"`
	Transactions   []Transaction `json:"transactions,omitempty" gorm:"foreignKey:AccountID"`
}
