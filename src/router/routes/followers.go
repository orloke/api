package routes

import (
	"api/src/controllers"
	"net/http"
)

func GetFollowersRoutes(c *controllers.FollowersController) []Route {
	return []Route{
		{
			Uri:    "/users/{id}/follow",
			Method: http.MethodPost,
			Func:   c.FollowUser,
			Auth:   true,
		},
		{
			Uri:    "/users/{id}/unfollow",
			Method: http.MethodPost,
			Func:   c.UnfollowUser,
			Auth:   true,
		},
		{
			Uri:    "/users/{id}/followers",
			Method: http.MethodGet,
			Func:   c.GetFollowers,
			Auth:   true,
		},
		{
			Uri:    "/users/{id}/following",
			Method: http.MethodGet,
			Func:   c.GetFollowing,
			Auth:   true,
		},
	}
}
