package main

import (
	"mugambi-ian/go-banking/handlers"
	"mugambi-ian/go-banking/logger"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting Application")
	sanityCheck()
	handlers.Start()
}

func sanityCheck() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal(err.Error())
	}
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_NAME") == "" ||
		os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PORT") == "" {
		logger.Fatal("Enviroment Variable Missing")
	}
}
