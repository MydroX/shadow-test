package main

import (
	"MydroX/shadow-technical-test/pkg/log"
	"MydroX/shadow-technical-test/src/server"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Initialize logger
	log.Logger = log.CreateLoggerDev()

	// Starting application
	log.Logger.Info("Starting application", zap.String("env", os.Getenv("APP_ENV")))
	server.Run()
}
