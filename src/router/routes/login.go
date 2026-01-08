package routes

import (
	"api/src/controllers"
	"net/http"
)

func GetLoginRoutes(c *controllers.LoginController) []Route {
	return []Route{
		{
			Uri:    "/login",
			Method: http.MethodPost,
			Func:   c.Login,
			Auth:   false,
		},
	}
}
