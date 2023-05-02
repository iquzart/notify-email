package main

import (
	"log"
	"notify-email/server"
	"os"
)

func main() {
	// Custom port
	port := ":" + os.Getenv("PORT")

	// Initialize the server
	server := server.InitServer()

	// Run server
	if port == ":" {
		log.Printf("Go App started on default port")
		server.Run()
	} else {
		log.Printf("Go App started on port " + os.Getenv("PORT"))
		server.Run(port)
	}

}
