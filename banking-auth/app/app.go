package app

import (
	"fmt"
	"go-banking/banking-auth/domain"
	"go-banking/banking-auth/handlers"
	"go-banking/banking-auth/service"
	"go-banking/banking-auth/utils"
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

	authRepository := domain.NewAuthRepository(getDBClient())
	authHandler := handlers.NewAuthHandler(service.NewLoginService(authRepository, domain.GetRolePermissions()))

	router.HandleFunc("/auth/login", authHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", authHandler.Register).Methods(http.MethodPost)
	router.HandleFunc("/auth/refresh", authHandler.Refresh).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", authHandler.Verify).Methods(http.MethodGet)

	port := os.Getenv("SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	utils.LogInfo(fmt.Sprintf("Starting OAuth server on %s:%s ...", address, port))
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
