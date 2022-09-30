package handlers

import (
	"log"
	"mugambi-ian/go-banking/domain"
	"mugambi-ian/go-banking/service"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {
	router := mux.NewRouter()
	dbClient := getDBClient()

	accountService := service.NewAccountService(domain.NewAccountRepositoryDB(dbClient))
	customerService := service.NewCustomerService(domain.NewCustomerRepositoryDB(dbClient))
	transactionService := service.NewTransactionService(domain.NewTransactionRepositoryDB(dbClient))

	accountHandler := AccountHandler{service: accountService}
	customerHandler := CustomerHandler{service: customerService}
	transactionHandler := TransactionHandler{service: transactionService}

	router.HandleFunc("/customers", customerHandler.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", customerHandler.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}/createAccount", accountHandler.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/transact/{account_id:[0-9]+}", transactionHandler.Transact).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getDBClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPassword := os.Getenv("DB_PASSWORD")

	if dbPassword != "" {
		dbUser += ":" + dbPassword
	}

	client, err := sqlx.Open("mysql", dbUser+"@tcp("+dbAddress+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
