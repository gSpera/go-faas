package main

import "net/http"

//Route is the route used to acces this function
var Route = "/example"

//Handle is the entrypoint to the function
func Handle(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is an example"))
}
