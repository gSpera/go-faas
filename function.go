package main

import (
	"fmt"
	"net/http"
	"plugin"
)

//functionHandler is the handler that a function musrt support.
type functionHandler = func(http.ResponseWriter, *http.Request)

//function contains information abount a function
type function struct {
	route   string
	handler functionHandler
}

//loadFunction loads a function
func loadFunction(path string) (function, error) {
	pl, err := plugin.Open(path)
	if err != nil {
		return function{}, fmt.Errorf("Cannot open %s plugin: %v", path, err)
	}
	routeRaw, err := pl.Lookup("Route")
	if err != nil {
		return function{}, fmt.Errorf("Cannot lookup Route: %v", err)
	}
	route := *routeRaw.(*string)

	handlerRaw, err := pl.Lookup("Handle")
	if err != nil {
		return function{}, fmt.Errorf("Cannot lookup handler: %v", err)
	}
	handler := handlerRaw.(functionHandler)

	return function{
		route:   route,
		handler: handler,
	}, nil
}
