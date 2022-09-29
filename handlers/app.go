package handlers

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
	router.HandleFunc("/customer/{customer_id:[0-9]+}", customerHandler.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
