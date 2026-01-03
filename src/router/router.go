package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	return routes.ConfigureRoutes(mux.NewRouter())
}
