package app

import (
	"fmt"
	"go-banking/banking-app/domain"
	"go-banking/banking-app/handlers"
	"go-banking/banking-app/middleware"
	"go-banking/banking-app/service"
	"go-banking/banking-app/utils"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {
	router := mux.NewRouter()
	dbClient := getDBClient()

	accountDb := domain.NewAccountRepositoryDB(dbClient)
	accountService := service.NewAccountService(accountDb)

	customerDb := domain.NewCustomerRepositoryDB(dbClient)
	customerService := service.NewCustomerService(customerDb)

	transactionDb := domain.NewTransactionRepositoryDB(dbClient)
	transactionService := service.NewTransactionService(transactionDb)

	accountHandler := handlers.NewAccountHandler(accountService)
	customerHandler := handlers.NewCustomerHandler(customerService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	router.HandleFunc("/customer", customerHandler.NewCustomer).
		Methods(http.MethodPost).
		Name("NewCustomer")

	router.HandleFunc("/customers", customerHandler.GetAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")

	router.HandleFunc("/customer/{customer_id:[0-9]+}", customerHandler.GetCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")

	router.HandleFunc("/transact/{account_id:[0-9]+}", transactionHandler.Transact).
		Methods(http.MethodPost).
		Name("NewTransaction")

	router.HandleFunc("/customer/{customer_id:[0-9]+}/createAccount", accountHandler.NewAccount).
		Methods(http.MethodPost).
		Name("NewAccount")

	authMiddleware := middleware.NewAuthMiddleware(domain.NewAuthRepository())
	router.Use(authMiddleware.AuthorizationHandler())

	port := os.Getenv("SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	utils.LogInfo(fmt.Sprintf("Starting banking server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
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
