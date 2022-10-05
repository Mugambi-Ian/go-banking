package domain

import (
	"go-banking/banking-app/dto"
	"go-banking/banking-app/utils"
)

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
	Account         Account
	Customer        Customer
}

type TransactionRepository interface {
	Transact(t Transaction) (*Transaction, *utils.AppError)
}

func (t Transaction) ToTransactionResponseDTO() dto.TransactionResponse {
	response := dto.TransactionResponse{
		Balance:       t.Account.Amount,
		TransactionId: t.TransactionId,
		Id:            t.Customer.Id,
		Name:          t.Customer.Name,
		AccountType:   t.Account.AccountType,
	}
	return response
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == "deposit" {
		return false
	} else {
		return true
	}
}
