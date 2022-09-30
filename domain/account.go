package domain

import (
	"mugambi-ian/go-banking/dto"
	"mugambi-ian/go-banking/errs"
)

type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

type AccountRespository interface {
	Save(Account) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}
