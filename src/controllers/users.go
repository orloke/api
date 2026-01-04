package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		responses.Error(w, http.StatusMethodNotAllowed, errors.New("Method not allowed"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, errors.New("Failed to read body request"))
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, errors.New("Error converting user to struct"))
		return
	}

	if err = user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, errors.New("Error connecting to database"))
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	insertedID, err := repository.Create(user)
	if err != nil {
		var driverErr *mysql.MySQLError
		if errors.As(err, &driverErr) && driverErr.Number == 1062 {
			if strings.Contains(driverErr.Message, "email") {
				responses.Error(w, http.StatusConflict, errors.New("email already registered"))
				return
			}
			if strings.Contains(driverErr.Message, "nick") {
				responses.Error(w, http.StatusConflict, errors.New("nick already registered"))
				return
			}
		}

		responses.Error(w, http.StatusInternalServerError, errors.New("Error creating user"))
		return
	}
	user.ID = int(insertedID)

	response := struct {
		Message string `json:"message"`
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Nick    string `json:"nick"`
		Email   string `json:"email"`
	}{
		Message: "User inserted",
		ID:      user.ID,
		Name:    user.Name,
		Nick:    user.Nick,
		Email:   user.Email,
	}

	responses.JSON(w, http.StatusCreated, response)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get Users"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get User"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete User"))
}
