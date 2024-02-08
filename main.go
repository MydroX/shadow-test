package main

import (
	"MydroX/shadow-technical-test/pkg/db/postgres"
	"MydroX/shadow-technical-test/pkg/log"
	"MydroX/shadow-technical-test/src/models"
	"MydroX/shadow-technical-test/src/server"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file ")
	}

	// Initialize logger
	log.Logger = log.CreateLoggerDev()

	// Initialize database and migrate models
	postgres.InitDB()
	err = postgres.Conn.AutoMigrate(&models.Team{}, &models.Player{})
	if err != nil {
		log.Logger.Fatal("Error migrating models", zap.Error(err))
	}

	// Starting application
	log.Logger.Info("Starting application", zap.String("env", os.Getenv("APP_ENV")))
	server.Run()
}
