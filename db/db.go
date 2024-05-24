package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost port=5432 dbname=events user=postgres password=postgres")
	if err != nil {
		log.Fatalf("Error connecting to the DB: v%\n", err)
	}

	defer db.Close()

	return db, nil
}
