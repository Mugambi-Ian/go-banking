package dto

import "go-banking/banking-app/utils"

type NewAccountRequest struct {
	CustomerId  string  `json:"CustomerId"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type NewAccountResponse struct {
	AccountId string `json:"account_id" `
}

func (nar NewAccountRequest) Validate() *utils.AppError {
	if nar.Amount < 5000 {
		return utils.NewValidationError("Low Deposit Amount")
	}
	if nar.AccountType != "saving" || nar.AccountType == "checking" {
		return utils.NewValidationError("Invalid Account Type")
	}
	return nil
}
