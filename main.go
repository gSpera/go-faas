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
	checkDebugFlag()

	if _, err := os.Stat(functionsDir); os.IsNotExist(err) {
		log("Creating %s directory\n", functionsDir)
		os.Mkdir(functionsDir, 0766)
	}

	fns, err := ioutil.ReadDir(functionsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read %s directory: %v\n", functionsDir, err)
		os.Exit(1)
	}

	for _, fn := range fns {
		if fn.IsDir() { //Sub-directory are not supported yet
			log("Skipping directory: %s\n", fn.Name())
			continue
		}

		log("Loading plugin: %s\n", fn.Name())
		plPath := path.Join(functionsDir, fn.Name())
		log(" -Path: %s\n", plPath)
		pl, err := loadFunction(plPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot load plugin %s: %v\n", fn.Name(), err)
		}

		log(" -Handling %s: %s\n", pl.route, fn.Name())
		http.HandleFunc(pl.route, pl.handler)
	}

	log("Listening: %s\n", addr)
	http.ListenAndServe(addr, nil)
}
