package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if error := authentication.ValidateToken(r); error != nil {
			responses.Error(w, http.StatusUnauthorized, error)
			return
		}

		next(w, r)
	}
}
