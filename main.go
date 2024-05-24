package main

import (
	"events-app/db"
	"log"
)

func main() {
	// connecting to the events DB
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("Connected to events DB")

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging DB: v%\n", err)
	}
	log.Println("Pinged database!")

}
