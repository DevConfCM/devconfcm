package router

import (
	"devconfcm/pkg/handlers"
	"net/http"
)

// add new routes here
var routes = []route{
	{"/api/users/:id", http.MethodGet, handlers.GetUser},
	{"/api/users", http.MethodPost, handlers.CreateUser},
	{"/api/users", http.MethodGet, handlers.GetAllUsers},
	{"/api/users/:id", http.MethodPut, handlers.UpdateUser},
	{"/api/users/:id", http.MethodDelete, handlers.DeleteUser},
}
