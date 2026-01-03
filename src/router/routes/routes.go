package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri    string
	Method string
	Func   http.HandlerFunc
	Auth   bool
}

func ConfigureRoutes(router *mux.Router) *mux.Router {
	for _, route := range UsersRoutes {
		router.HandleFunc(route.Uri, route.Func).Methods(route.Method)
	}

	return router
}
