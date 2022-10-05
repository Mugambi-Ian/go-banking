package main

import (
	"fmt"
	"go-banking/banking-auth/app"
	"go-banking/banking-auth/utils"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	sanityCheck()
	app.Start()
}

func sanityCheck() {
	err := godotenv.Load()
	if err != nil {
		utils.LogError(err.Error())
	}
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_ADDRESS",
		"DB_USER",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			utils.LogError(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}
