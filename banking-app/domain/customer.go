package domain

import (
	"go-banking/banking-app/dto"
	"go-banking/banking-app/utils"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	Create(Customer) (*Customer, *utils.AppError)
	FindAll(string) ([]Customer, *utils.AppError)
	ByID(string) (*Customer, *utils.AppError)
}

func (c Customer) ToDto() dto.CustomerResponse {
	response := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		DateofBirth: c.DateofBirth,
		Zipcode:     c.Zipcode,
		Status:      c.statusAsText(),
		City:        c.City,
	}
	return response
}

func (c Customer) statusAsText() string {
	statusText := "active"
	if c.Status == "0" {
		statusText = "inactive"
	}
	return statusText
}
