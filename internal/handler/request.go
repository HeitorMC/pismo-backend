package handler

import (
	"fmt"
	"regexp"
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
