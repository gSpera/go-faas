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

	fns := index(functionsDir)

	for _, fn := range fns {
		log("Handling %s\n", fn.route)
		http.HandleFunc(fn.route, fn.handler)
	}

	log("Listening: %s\n", addr)
	http.ListenAndServe(addr, nil)
}

//index indexes the folder given as parameter and returns all the functions contained inside the folder
func index(folder string) []function {
	var functions []function

	log("Indexing: %s\n", folder)

	fns, err := ioutil.ReadDir(folder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read %s directory: %v\n", functionsDir, err)
		os.Exit(1)
	}

	functions = make([]function, 0, len(fns))

	for _, fn := range fns {
		if fn.IsDir() { //Sub-directory are not supported yet
			log(" -Skipping directory: %s\n", fn.Name())
			continue
		}

		log("Loading plugin: %s\n", fn.Name())
		plPath := path.Join(folder, fn.Name())
		log("  -Path: %s\n", plPath)
		pl, err := loadFunction(plPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot load plugin %s: %v\n", fn.Name(), err)
		}

		functions = append(functions, pl)
	}

	return functions
}
