package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	DB, err := sql.Open("pgx", "host=localhost port=5432 dbname=events user=postgres password=postgres")
	if err != nil {
		log.Fatalf("Error connecting to the DB: v%\n", err)
	}

	return DB, nil
}
