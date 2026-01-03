package routes

import (
	"api/src/controllers"
	"net/http"
)

var UsersRoutes = []Route{
	{
		Uri:    "/users",
		Method: http.MethodPost,
		Func:   controllers.CreateUser,
		Auth:   false,
	},
	{
		Uri:    "/users",
		Method: http.MethodGet,
		Func:   controllers.GetUsers,
		Auth:   false,
	},
	{
		Uri:    "/users/{id}",
		Method: http.MethodGet,
		Func:   controllers.GetUser,
		Auth:   false,
	},
	{
		Uri:    "/users/{id}",
		Method: http.MethodPut,
		Func:   controllers.UpdateUser,
		Auth:   false,
	},
	{
		Uri:    "/users/{id}",
		Method: http.MethodDelete,
		Func:   controllers.DeleteUser,
		Auth:   false,
	},
}