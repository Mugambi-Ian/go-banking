package handlers

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"mugambi-ian/go-banking/dto"
	"mugambi-ian/go-banking/errs"
	"mugambi-ian/go-banking/service"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (customerHandler *CustomerHandler) getAllCustomers(res http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	status := queryParams.Get("status")
	customers := make([]dto.CustomerResponse, 0)
	err := errs.NewUnexpectedError("A Error Occurred")
	if status != "" {
		switch status {
		case "active":
			customers, err = customerHandler.service.GetAllCustomer("1")
			break
		case "inactive":
			customers, err = customerHandler.service.GetAllCustomer("0")
			break
		default:
			err = errs.NewUnexpectedError("Unrecognized Status Value")
			break
		}
	} else {
		customers, err = customerHandler.service.GetAllCustomer("")
	}
	if err != nil {
		sendJSONResponse(res, err.Code, err.GetMessage())
	} else if req.Header.Get("Content-Type") == "application/xml" {
		sendXMLResponse(res, 200, customers)
	} else {
		sendJSONResponse(res, 200, customers)
	}
}

func (customerHandler *CustomerHandler) getCustomer(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["customer_id"]
	customer, err := customerHandler.service.GetCustomer(id)
	if err != nil {
		sendJSONResponse(res, err.Code, err.GetMessage())
	} else if req.Header.Get("Content-Type") == "application/xml" {
		sendXMLResponse(res, 200, customer)
	} else {
		sendJSONResponse(res, 200, customer)
	}
}

func sendJSONResponse(res http.ResponseWriter, code int, data interface{}) {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(code)
	if err := json.NewEncoder(res).Encode(data); err != nil {
		panic(err.Error())
	}
}

func sendXMLResponse(res http.ResponseWriter, code int, data interface{}) {
	res.Header().Add("Content-Type", "application/xml")
	res.WriteHeader(code)
	//fmt.Fprint(res, err.Message)
	if err := xml.NewEncoder(res).Encode(data); err != nil {
		panic(err.Error())
	}
}
