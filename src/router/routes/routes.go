package routes

import (
	"api/src/controllers"
	"api/src/middlewares"
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
	loginController := controllers.NewLoginController(db)
	followersController := controllers.NewFollowersController(db)

	routes := GetUsersRoutes(usersController)
	routes = append(routes, GetLoginRoutes(loginController)...)
	routes = append(routes, GetFollowersRoutes(followersController)...)

	for _, route := range routes {
		if route.Auth {
			router.HandleFunc(route.Uri, middlewares.Authenticate(route.Func)).Methods(route.Method)
		} else {
			router.HandleFunc(route.Uri, route.Func).Methods(route.Method)
		}
	}

	return router
}
