package dto

import (
	"database/sql"
	"go-banking/banking-auth/domain"
	"go-banking/banking-auth/utils"
)

type RegisterRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	CustomerID string `json:"customer_id"`
}

func (r RegisterRequest) IsValid() *utils.AppError {
	if r.Username == "" {
		return utils.NewBadRequestError("Username Is Required")
	} else if r.Password == "" {
		return utils.NewBadRequestError("Password Is Required")
	} else if len(r.Password) < 8 {
		return utils.NewBadRequestError("Password Is Too Short")
	} else if r.Role != "admin" && r.Role != "user" {
		return utils.NewBadRequestError("Invalid Role")
	} else if r.Role == "user" && r.CustomerID == "" {
		return utils.NewBadRequestError("Customer Id Required")
	}
	return nil
}

func (r RegisterRequest) ToLoginCredentials() (*domain.Login, *utils.AppError) {
	password, err := utils.HashPassword(r.Password)
	if err != nil {
		utils.LogError(err.Message)
		return nil, err
	}
	utils.LogInfo(password)
	return &domain.Login{
		Username:   r.Username,
		Password:   password,
		CustomerId: sql.NullString{String: r.CustomerID, Valid: true},
		Role:       r.Role,
	}, nil
}
