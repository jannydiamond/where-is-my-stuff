package dbroutes

import (
	"database/sql"
	"log"

	"wims/internal/models"
)

func InsertUser(db *sql.DB, username string, password string) {
	insertSQL := "INSERT INTO wims_user (username, password) VALUES ($1, $2)"
	_, err := db.Exec(insertSQL, username, password)
	if err != nil {
		log.Fatal(err)
	}
}

func ListUsers(db *sql.DB) []models.UserResponse {
	selectSQL := "SELECT id, username, created_at, updated_at FROM wims_user"
	rows, err := db.Query(selectSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []models.UserResponse

	for rows.Next() {
		var u models.UserResponse
		if err := rows.Scan(&u.ID, &u.Username, &u.CreatedAt, &u.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

func UpdateUser(db *sql.DB, id string, password string) {
	updateSQL := "UPDATE wims_user SET password = $2, updated_at = NOW() WHERE id = $1"
	_, err := db.Exec(updateSQL, id, password)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUser(db *sql.DB, id string) models.UserResponse {
	selectSQL := "SELECT id, username, created_at, updated_at FROM wims_user WHERE id = $1"
	row := db.QueryRow(selectSQL, id)

	var u models.UserResponse
	err := row.Scan(&u.ID, &u.Username, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		log.Fatal(err)
	}

	return u
}

func DeleteUser(db *sql.DB, id string) {
	deleteSQL := "DELETE FROM wims_user WHERE id = $1"
	_, err := db.Exec(deleteSQL, id)
	if err != nil {
		log.Fatal(err)
	}
}
