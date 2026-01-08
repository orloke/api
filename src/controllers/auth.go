package controllers

import (
	"api/src/authentication"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type LoginController struct {
	DB *sql.DB
}

func NewLoginController(db *sql.DB) *LoginController {
	return &LoginController{DB: db}
}

func (l *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var loginData models.LoginData
	if err = json.Unmarshal(body, &loginData); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = loginData.Validate(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	repository := repositories.NewUsersRepository(l.DB)
	userInDB, err := repository.FindByEmailOrNick(loginData.User)
	if err != nil {
		if err == sql.ErrNoRows {
			responses.Error(w, http.StatusUnauthorized, errors.New("User or password invalid"))
			return
		}
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(userInDB.Password, loginData.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("User or password invalid"))
		return
	}

	token, err := authentication.CreateToken(uint64(userInDB.ID), userInDB.Email, userInDB.Nick)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, models.TokenResponse{
		Token: token,
	})
}