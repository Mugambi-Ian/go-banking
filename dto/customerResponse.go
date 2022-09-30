package dto

type CustomerResponse struct {
	Id          string `json:"customer_id" xml:"CustomerId"`
	Name        string `json:"customer_name" xml:"CustomerName"`
	City        string `json:"customer_city" xml:"CustomerCity"`
	Zipcode     string `json:"customer_zipcode" xml:"CustomerZipcode"`
	DateofBirth string `json:"customer_dateOfBirth" xml:"CustomerDateOfBirth"`
	Status      string `json:"customer_status" xml:"CustomerStatus"`
}
