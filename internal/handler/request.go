package handler

import (
	"fmt"
	"regexp"
	"strings"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

func (r CreateAccountRequest) Validate() error {
	if r.DocumentNumber == "" {
		return errParamIsRequired("document_number", "string")
	}
	if len(r.DocumentNumber) != 11 {
		return fmt.Errorf("param: document_number (type: string) must be 11 characters long")
	}

	re := regexp.MustCompile(`^\d{11}$`)
	if !re.MatchString(r.DocumentNumber) {
		return fmt.Errorf("param: document_number (type: string) must be 11 numeric characters long")
	}
	return nil
}

type CreateTransactionRequest struct {
	AccountID       *uint    `json:"account_id"`
	OperationTypeID *uint    `json:"operation_type_id"`
	Amount          *float64 `json:"amount"`
}

func (r CreateTransactionRequest) Validate() error {
	var errs []string

	if r.AccountID == nil {
		errs = append(errs, errParamIsRequired("account_id", "uint").Error())
	}
	if r.OperationTypeID == nil {
		errs = append(errs, errParamIsRequired("operation_type_id", "uint").Error())
	}
	if r.Amount == nil {
		errs = append(errs, errParamIsRequired("amount", "float64").Error())
	}

	if r.AccountID != nil && *r.AccountID == 0 {
		errs = append(errs, "param: account_id (type: uint) must be a non-zero value")
	}
	if r.OperationTypeID != nil && *r.OperationTypeID == 0 {
		errs = append(errs, "param: operation_type_id (type: uint) must be a non-zero value")
	}
	if r.Amount != nil && *r.Amount == 0 {
		errs = append(errs, "param: amount (type: float64) must be a non-zero value")
	}

	if len(errs) == 0 {
		return nil
	}
	return fmt.Errorf("validation errors: %s", strings.Join(errs, ", "))
}
