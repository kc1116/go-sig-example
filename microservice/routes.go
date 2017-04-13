package main

import (
	"net/http"
)

//Route . . .
//struct used to define a route,
//Name: description of the route
//Method: http method
//Pattern: actual route endpoint
//HandlerFunc: function to handle the route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//HealthCheckUp . . . struct should be returned from /health route with UP status.
type HealthCheckUp struct {
	Status      string `json:"status"`
	Application ApplicationStatus
}
type ApplicationStatus struct {
	Status string `json:"status"`
	Name   string `json:"name"`
}

//Routes slice of routes that are registered with the mux router. All new routes can be defined here.
type Routes []Route

var routes = Routes{
	Route{
		"Return microservice health status.", "GET", "/health", health,
	},
}
