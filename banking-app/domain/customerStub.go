package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "John A. Doe", City: "Nairobi", Zipcode: "00100", DateofBirth: "00/00/0000", Status: "1"},
		{Id: "1002", Name: "John B. Doe", City: "Nairobi", Zipcode: "00100", DateofBirth: "00/00/0000", Status: "1"},
		{Id: "1003", Name: "John C. Doe", City: "Nairobi", Zipcode: "00100", DateofBirth: "00/00/0000", Status: "1"},
		{Id: "1004", Name: "John D. Doe", City: "Nairobi", Zipcode: "00100", DateofBirth: "00/00/0000", Status: "1"},
		{Id: "1005", Name: "John E. Doe", City: "Nairobi", Zipcode: "00100", DateofBirth: "00/00/0000", Status: "1"},
		{Id: "1006", Name: "John F. Doe", City: "Nairobi", Zipcode: "00100", DateofBirth: "00/00/0000", Status: "1"},
		{Id: "1007", Name: "John C. Doe", City: "Nairobi", Zipcode: "00100", DateofBirth: "00/00/0000", Status: "1"}}
	return CustomerRepositoryStub{customers: customers}
}
