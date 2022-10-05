package handlers

import (
	"encoding/json"
	"go-banking/banking-app/dto"
	"go-banking/banking-app/service"
	"go-banking/banking-app/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			utils.SendJSONResponse(res, appError.Code, appError.Message)
		} else {
			utils.SendJSONResponse(res, http.StatusCreated, account)
		}
	}
}

func NewAccountHandler(s service.AccountService) AccountHandler {
	return AccountHandler{service: s}
}
