package main

import (
	"fmt"
	"net/http"
	"plugin"
)

func main() {
	fn, err := plugin.Open("example.so")
	if err != nil {
		panic(err)
	}
	routeRaw, err := fn.Lookup("Route")
	if err != nil {
		panic(err)
	}
	route := *routeRaw.(*string)
	fmt.Println(route)
	handlerRaw, err := fn.Lookup("Handle")
	if err != nil {
		panic(err)
	}
	handler := handlerRaw.(func(http.ResponseWriter, *http.Request))
	http.HandleFunc(route, handler)
	http.ListenAndServe(":8080", nil)
}
