package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(path_prefix + "css/"))))
	http.Handle("/css", router)

	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(path_prefix + "js/"))))
	http.Handle("/js", router)

	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir(path_prefix + "img/"))))
	http.Handle("/img", router)
	return router
}
