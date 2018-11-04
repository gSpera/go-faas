# Go Faas
## WIP: This is a WIP project

Go Faas is an attempt to port faas(function as a service) to a simpler level,
this can be accomplished thanks to Golang plugin support(currently not for Windows).

## Creating a Function
The main goal is that writing a function should be the simplest possible.
A function is made of a path (*Route*) and an handler(*Handler*)
```go
package main

import "net/http"

//Route is the route used to acces this function
var Route = "/example"

//Handle is the entrypoint to the function
func Handle(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is an example"))
}

```
Building is achived thanks to plugins
`go build --buildmode=plugin example.go`

Now copy the output to the functions folder within the go-faas executable
`cp example /path/to/go-faas/functions/`

The last step is to notify the go-faas process to search for new functions, this can be achived restarting the process
or sending an USR1 signal
`kill -s USR1 $(cat /path/to/go-faas/pid.pid)`

As you can see the main focus is to be runned over an web-server, but also a cli interface is avaible throught
`/path/to/go-faas example`
Specify the stdin flag to send data using POST body and use parameter in the form `key=value` to send GET parameter.