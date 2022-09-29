package app

import (
	"log"
	"mugambi-ian/go-banking/domain"
	"mugambi-ian/go-banking/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	customerService := service.NewCustomerService(domain.NewCustomerRepositoryDB())
	customerHandler := CustomerHandler{service: customerService}

	router.HandleFunc("/customers", customerHandler.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
