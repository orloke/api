package routes

import (
	"api/src/controllers"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri    string
	Method string
	Func   http.HandlerFunc
	Auth   bool
}

func ConfigureRoutes(router *mux.Router, db *sql.DB) *mux.Router {
	usersController := controllers.NewUsersController(db)

	for _, route := range GetUsersRoutes(usersController) {
		router.HandleFunc(route.Uri, route.Func).Methods(route.Method)
	}

	return router
}
