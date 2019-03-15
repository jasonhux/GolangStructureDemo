package main

import (
	"GolangStructureDemo/src/common"
	"GolangStructureDemo/src/endpoints/home"
	"GolangStructureDemo/src/endpoints/search"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := getRoutes()

	for _, route := range *routes {
		router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Path).
			HandlerFunc(route.Handler)
	}
	return router
}

func getRoutes() *Routes {
	ss := search.NewService([]string{"Australia", "China", "New Zealand", "America", "England"})
	return &Routes{
		Route{
			"Home",
			"GET",
			"/",
			common.GetHandler(home.NewEndpoint()),
		},
		Route{
			"Search",
			"GET",
			"/search",
			common.GetHandler(search.NewEndpoint(ss)),
		},
	}
}
