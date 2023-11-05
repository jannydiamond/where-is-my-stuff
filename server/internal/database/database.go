package database

import (
	"database/sql"
	"fmt"
	"os"
)

func Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("PGUSER"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGPASSWORD"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
