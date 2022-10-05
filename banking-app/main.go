package main

import (
	"fmt"
	"go-banking/banking-app/app"
	"os"

	"github.com/joho/godotenv"
	"go-banking/banking-app/utils"
)

func main() {
	sanityCheck()
	app.Start()
}

func sanityCheck() {
	err := godotenv.Load()
	if err != nil {
		utils.LogFatal(err.Error())
	}
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_ADDRESS",
		"DB_USER",
		"DB_PORT",
		"DB_NAME",
		"AUTH_URI",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			utils.LogFatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}
