package domain

import (
	"mugambi-ian/go-banking/dto"
	"mugambi-ian/go-banking/errs"
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
	FindAll(string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}

func (c Customer) ToDto() dto.CustomerResponse {
	response := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		DateofBirth: c.DateofBirth,
		Zipcode:     c.Zipcode,
		Status:      c.statusAsText(),
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
