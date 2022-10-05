package handlers

import (
	"encoding/json"
	"go-banking/banking-app/dto"
	"go-banking/banking-app/service"
	"go-banking/banking-app/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (customerHandler CustomerHandler) GetAllCustomers(res http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	status := queryParams.Get("status")
	customers := make([]dto.CustomerResponse, 0)
	err := utils.NewUnexpectedError("A Error Occurred")
	if status != "" {
		switch status {
		case "active":
			customers, err = customerHandler.service.GetAllCustomer("1")
			break
		case "inactive":
			customers, err = customerHandler.service.GetAllCustomer("0")
			break
		default:
			err = utils.NewUnexpectedError("Unrecognized Status Value")
			break
		}
	} else {
		customers, err = customerHandler.service.GetAllCustomer("")
	}
	if err != nil {
		utils.SendJSONResponse(res, err.Code, err.GetMessage())
	} else if req.Header.Get("Content-Type") == "application/xml" {
		utils.SendXMLResponse(res, 200, customers)
	} else {
		utils.SendJSONResponse(res, 200, customers)
	}
}

func (customerHandler CustomerHandler) GetCustomer(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["customer_id"]
	customer, err := customerHandler.service.GetCustomer(id)
	if err != nil {
		utils.SendJSONResponse(res, err.Code, err.GetMessage())
	} else if req.Header.Get("Content-Type") == "application/xml" {
		utils.SendXMLResponse(res, 200, customer)
	} else {
		utils.SendJSONResponse(res, 200, customer)
	}
}

func (customerHandler CustomerHandler) NewCustomer(res http.ResponseWriter, req *http.Request) {
	var request dto.NewCustomerRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, err.Error())
	} else {
		account, appError := customerHandler.service.NewCustomer(request)
		if appError != nil {
			utils.SendJSONResponse(res, appError.Code, appError.Message)
		} else {
			utils.SendJSONResponse(res, http.StatusCreated, account)
		}
	}
}

func NewCustomerHandler(s service.CustomerService) CustomerHandler {
	return CustomerHandler{
		service: s,
	}
}
