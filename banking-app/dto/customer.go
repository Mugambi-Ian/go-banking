package dto

import (
	"go-banking/banking-app/utils"
)

type CustomerResponse struct {
	Id          string `json:"customer_id" xml:"CustomerId"`
	Name        string `json:"customer_name" xml:"CustomerName"`
	City        string `json:"customer_city" xml:"CustomerCity"`
	Zipcode     string `json:"customer_zipcode" xml:"CustomerZipcode"`
	DateofBirth string `json:"customer_dateOfBirth" xml:"CustomerDateOfBirth"`
	Status      string `json:"customer_status" xml:"CustomerStatus"`
}

type NewCustomerRequest struct {
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `json:"date_of_birth"`
}

func (r NewCustomerRequest) IsValid() *utils.AppError {
	if r.Name == "" {
		return utils.NewBadRequestError("Name is Required")
	} else if r.City == "" {
		return utils.NewBadRequestError("City is Required")
	} else if r.Zipcode == "" {
		return utils.NewBadRequestError("Zipcode is Required")
	} else if r.DateofBirth == "" {
		return utils.NewBadRequestError("Invalid date of birth")
	}
	return nil
}
