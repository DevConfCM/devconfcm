package routes

import (
	"devconfcm/pkg/handlers"
	"net/http"
)

// add new routes here
var routes = []Route{
	{"/api/user/:id/", http.MethodGet, handlers.GetUser},
	{"/api/user/", http.MethodPost, handlers.CreateUser},
	{"/api/users/", http.MethodGet, handlers.GetAllUsers},
}
