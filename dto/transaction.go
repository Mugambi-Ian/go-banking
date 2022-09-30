package dto

import "mugambi-ian/go-banking/errs"

type TransactionRequest struct {
	AccountId       string
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (t TransactionRequest) Validate() *errs.AppError {
	if t.Amount < 1 {
		return errs.NewValidationError("Invalid Transaction Amount")
	}
	if t.TransactionType != "withdraw" && t.TransactionType != "deposit" {
		return errs.NewValidationError("Invalid Transaction Type")
	}
	return nil
}

type TransactionResponse struct {
	Balance       float64
	TransactionId string `json:"transaction_id"`
}
