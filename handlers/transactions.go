package handlers

import (
	"encoding/json"
	"mugambi-ian/go-banking/dto"
	"mugambi-ian/go-banking/service"
	"net/http"

	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (h TransactionHandler) Transact(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	accountId := vars["account_id"]
	var request dto.TransactionRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		sendJSONResponse(res, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		transaction, appError := h.service.Transact(request)
		if appError != nil {
			sendJSONResponse(res, appError.Code, appError.Message)
		} else {
			sendJSONResponse(res, http.StatusCreated, transaction)
		}
	}
}
