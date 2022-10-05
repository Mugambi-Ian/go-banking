package domain

import (
	"go-banking/banking-app/dto"
	"go-banking/banking-app/utils"
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
	Save(Account) (*Account, *utils.AppError)
}

func (a Account) ToNewAccountResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}
