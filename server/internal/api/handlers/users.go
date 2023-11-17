package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"wims/internal/dbroutes"
	"wims/internal/models"
	"wims/internal/utils"

	"github.com/go-chi/chi/v5"
)

func ListUsersHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	users := dbroutes.ListUsers(db)

	utils.RespondWithJSON(w, http.StatusOK, users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req models.User

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error decoding request body")
	}

	dbroutes.InsertUser(db, req.Username, req.Password)

	utils.RespondWithJSON(w, http.StatusOK, "Created successfully")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := chi.URLParam(r, "userId")
	user := dbroutes.GetUser(db, id)

	utils.RespondWithJSON(w, http.StatusOK, user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var req models.User

	id := chi.URLParam(r, "userId")
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error decoding request body")
	}

	dbroutes.UpdateUser(db, id, req.Password)

	user := dbroutes.GetUser(db, id)
	utils.RespondWithJSON(w, http.StatusOK, user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := chi.URLParam(r, "userId")

	dbroutes.DeleteUser(db, id)

	utils.RespondWithJSON(w, http.StatusOK, "Deleted successfully")
}
