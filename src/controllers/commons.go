package controllers

import (
	"api/src/responses"
	"errors"
	"net/http"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func handleDuplicateError(w http.ResponseWriter, err error) bool {
	var driverErr *mysql.MySQLError
	if errors.As(err, &driverErr) {
		if driverErr.Number == 1062 {
			if strings.Contains(driverErr.Message, "email") {
				responses.Error(w, http.StatusConflict, errors.New("email already registered"))
				return true
			}
			if strings.Contains(driverErr.Message, "nick") {
				responses.Error(w, http.StatusConflict, errors.New("nick already registered"))
				return true
			}
			if strings.Contains(driverErr.Message, "followers") {
				responses.Error(w, http.StatusForbidden, errors.New("User already following"))
				return true
			}
		}

		if driverErr.Number == 1452 {
			responses.Error(w, http.StatusNotFound, errors.New("User not found"))
			return true
		}
	}
	return false
}
