package domain

import (
	"mugambi-ian/go-banking/dto"
	"mugambi-ian/go-banking/errs"
)

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
	Balance         float64
}

type TransactionRepository interface {
	Transact(t Transaction) (*Transaction, *errs.AppError)
}

func (t Transaction) ToTransactionResponseDTO(amount float64) dto.TransactionResponse {
	response := dto.TransactionResponse{
		Balance:       amount,
		TransactionId: t.TransactionId,
	}
	return response
}
