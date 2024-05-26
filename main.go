package main

import (
	"events-app/db"
	"events-app/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// connecting to the events DB
	dtbs, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer dtbs.Close()

	// create HTTP server
	server := gin.Default()

	// HTTP routes from routes.go
	routes.RegisterRoutes(server)

	// start server at localhost:8080
	err = server.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: v%\n", err)
	}
}
