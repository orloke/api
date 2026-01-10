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
			Auth:   true,
		},
		{
			Uri:    "/users",
			Method: http.MethodGet,
			Func:   c.GetUsers,
			Auth:   true,
		},
		{
			Uri:    "/users/{id}",
			Method: http.MethodGet,
			Func:   c.GetUser,
			Auth:   true,
		},
		{
			Uri:    "/users/{id}",
			Method: http.MethodPut,
			Func:   c.UpdateUser,
			Auth:   true,
		},
		{
			Uri:    "/users/{id}",
			Method: http.MethodDelete,
			Func:   c.DeleteUser,
			Auth:   true,
		},
		{
			Uri:    "/users/{id}/update-password",
			Method: http.MethodPost,
			Func:   c.UpdatePassword,
			Auth:   true,
		},
	}
}
