package controllers

import (
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type UsersController struct {
	DB *sql.DB
}

func NewUsersController(db *sql.DB) *UsersController {
	return &UsersController{DB: db}
}

func (c *UsersController) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if err = user.Prepare("register"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	repository := repositories.NewUsersRepository(c.DB)
	insertedID, err := repository.Create(user)
	if err != nil {
		if handleDuplicateError(w, err) {
			return
		}

		responses.Error(w, http.StatusInternalServerError, errors.New("Error creating user"))
		return
	}
	user.ID = int(insertedID)

	response := models.CreateUserResponse{
		Message: "User inserted",
		User: models.NewUserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Nick:  user.Nick,
			Email: user.Email,
		},
	}

	responses.JSON(w, http.StatusCreated, response)
}

func (c *UsersController) GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	repository := repositories.NewUsersRepository(c.DB)
	users, err := repository.Search(nameOrNick)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func (c *UsersController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, errors.New("Invalid user ID"))
		return
	}

	repository := repositories.NewUsersRepository(c.DB)
	user, err := repository.FindByID(userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			responses.Error(w, http.StatusNotFound, errors.New("User not found"))
			return
		}
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func (c *UsersController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, errors.New("Invalid user ID"))
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

	repository := repositories.NewUsersRepository(c.DB)
	userInDB, err := repository.FindByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			responses.Error(w, http.StatusNotFound, errors.New("User not found"))
			return
		}
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if user.Name == "" {
		user.Name = userInDB.Name
	}
	if user.Nick == "" {
		user.Nick = userInDB.Nick
	}
	if user.Email == "" {
		user.Email = userInDB.Email
	}

	if err = user.Prepare("edit"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.Update(userID, user); err != nil {
		if handleDuplicateError(w, err) {
			return
		}
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func (c *UsersController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, errors.New("Invalid user ID"))
		return
	}

	repository := repositories.NewUsersRepository(c.DB)
	if err = repository.Delete(userID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
