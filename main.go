package main

import (
	"events-app/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// connecting to the events DB
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// create HTTP server
	server := gin.Default()

	// start server at localhost:8080
	err = server.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: v%\n", err)
	}
}
