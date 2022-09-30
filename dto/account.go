package dto

import "mugambi-ian/go-banking/errs"

type NewAccountRequest struct {
	CustomerId  string  `json:"CustomerId"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type NewAccountResponse struct {
	AccountId string `json:"account_id" `
}

func (nar NewAccountRequest) Validate() *errs.AppError {
	if nar.Amount < 5000 {
		return errs.NewValidationError("Low Deposit Amount")
	}
	if nar.AccountType != "saving" || nar.AccountType == "checking" {
		return errs.NewValidationError("Invalid Account Type")
	}
	return nil
}
