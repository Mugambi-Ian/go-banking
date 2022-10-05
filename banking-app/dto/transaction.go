package dto

import "go-banking/banking-app/utils"

type TransactionRequest struct {
	AccountId       string
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (t TransactionRequest) Validate() *utils.AppError {
	if t.Amount < 1 {
		return utils.NewValidationError("Invalid Transaction Amount")
	}
	if t.TransactionType != "withdraw" && t.TransactionType != "deposit" {
		return utils.NewValidationError("Invalid Transaction Type")
	}
	return nil
}

type TransactionResponse struct {
	Id            string  `json:"customer_id"`
	Name          string  `json:"customer_name"`
	TransactionId string  `json:"transaction_id"`
	Balance       float64 `json:"account_balance"`
	AccountType   string  `db:"account_type"`
}
