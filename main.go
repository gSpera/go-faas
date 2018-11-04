package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

const (
	functionsDir = "functions"
	addr         = ":7070"
)

func main() {
	fns, err := ioutil.ReadDir(functionsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read %s directory: %v\n", functionsDir, err)
		os.Exit(1)
	}

	for _, fn := range fns {
		if fn.IsDir() { //Sub-directory are not supported yet
			continue
		}

		plPath := path.Join(functionsDir, fn.Name())
		pl, err := loadFunction(plPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot load plugin %s: %v\n", fn.Name(), err)
		}
		http.HandleFunc(pl.route, pl.handler)
	}

	http.ListenAndServe(addr, nil)
}
