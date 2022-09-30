package handlers

import (
	"encoding/json"
	"mugambi-ian/go-banking/dto"
	"mugambi-ian/go-banking/service"
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
		sendJSONResponse(res, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			sendJSONResponse(res, appError.Code, appError.Message)
		} else {
			sendJSONResponse(res, http.StatusCreated, account)
		}
	}
}
