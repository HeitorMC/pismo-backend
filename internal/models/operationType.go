package models

import "gorm.io/gorm"

type OperationType struct {
	OperationTypeID uint          `json:"operation_type_id" gorm:"primaryKey"`
	Description     string        `json:"description" gorm:"unique;not null"`
	Transactions    []Transaction `json:"transactions" gorm:"foreignKey:OperationTypeID"`
}

// Seed operation types
func SeedOperationTypes(db *gorm.DB) {
	operationTypes := []OperationType{
		{OperationTypeID: 1, Description: "COMPRA A VISTA"},
		{OperationTypeID: 2, Description: "COMPRA PARCELADA"},
		{OperationTypeID: 3, Description: "SAQUE"},
		{OperationTypeID: 4, Description: "PAGAMENTO"},
	}

	for _, op := range operationTypes {
		db.FirstOrCreate(&op, op)
	}
}
