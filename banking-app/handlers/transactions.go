package handlers

import (
	"encoding/json"
	"go-banking/banking-app/dto"
	"go-banking/banking-app/service"
	"go-banking/banking-app/utils"
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
		utils.SendJSONResponse(res, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		transaction, appError := h.service.Transact(request)
		if appError != nil {
			utils.SendJSONResponse(res, appError.Code, appError.Message)
		} else {
			utils.SendJSONResponse(res, http.StatusCreated, transaction)
		}
	}
}

func NewTransactionHandler(s service.TransactionService) TransactionHandler {
	return TransactionHandler{
		service: s,
	}
}
