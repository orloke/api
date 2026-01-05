package router

import (
	"api/src/router/routes"
	"database/sql"

	"github.com/gorilla/mux"
)

func Router(db *sql.DB) *mux.Router {
	return routes.ConfigureRoutes(mux.NewRouter(), db)
}
