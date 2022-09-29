package app

import (
	"encoding/json"
	"encoding/xml"
	"mugambi-ian/go-banking/service"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (customerHandler *CustomerHandler) getAllCustomers(res http.ResponseWriter, req *http.Request) {
	customers, _ := customerHandler.service.GetAllCustomer()

	if req.Header.Get("Content-Type") == "application/xml" {
		res.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(res).Encode(customers)
	} else {
		res.Header().Add("Content-Type", "application/json")
		json.NewEncoder(res).Encode(customers)
	}
}
