package main

import "log"

func init() {

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
