package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user='%s' password='%s' dbname='%s' sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("SCHEMA"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
