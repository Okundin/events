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

}
