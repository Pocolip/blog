package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		indexHandler,
	},
	Route{
		"BlogShow",
		"GET",
		"/blog/{blogId}/",
		viewHandler,
	},
	Route{ //Unused
		"BlogCreate",
		"POST",
		"/blog/",
		createHandler,
	},
}
