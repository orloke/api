package routes

import (
	"api/src/controllers"
	"net/http"
)

func GetUsersRoutes(c *controllers.UsersController) []Route {
	return []Route{
		{
			Uri:    "/users",
			Method: http.MethodPost,
			Func:   c.CreateUser,
			Auth:   false,
		},
		{
			Uri:    "/users",
			Method: http.MethodGet,
			Func:   c.GetUsers,
			Auth:   false,
		},
		{
			Uri:    "/users/{id}",
			Method: http.MethodGet,
			Func:   c.GetUser,
			Auth:   false,
		},
		{
			Uri:    "/users/{id}",
			Method: http.MethodPut,
			Func:   c.UpdateUser,
			Auth:   false,
		},
		{
			Uri:    "/users/{id}",
			Method: http.MethodDelete,
			Func:   c.DeleteUser,
			Auth:   false,
		},
	}
}
