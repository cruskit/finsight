package restservice

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
		"/finsight/",
		Index,
	},
	Route{
		"Index",
		"GET",
		"/finsight/strategy",
		RunStrategy,
	},
}
