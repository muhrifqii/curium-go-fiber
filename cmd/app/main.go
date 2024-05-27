package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	server, err := InitializeServer()
	if err != nil {
		log.Fatal("Failed to initialize server: %v", err)
	}
	if err := server.App.Listen(":3000"); err != nil {
		log.Fatal("Failed to start server: %v", err)
	}
}
