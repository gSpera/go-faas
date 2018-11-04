package main

import (
	"fmt"
	"os"
)

//debug.go contains function definition for debug functions

//debugFlag is set when the debug functions are enabled
var debugFlag bool

const debugEnvVar = "GO_FAAS_DEBUG"

//checkDebugFlag sets the debugFlag variable if the enviroment conditions are set.
//Condition:
//GO_FAAS_DEBUG=ENABLE_DEBUG
func checkDebugFlag() {
	value := os.Getenv(debugEnvVar)
	if value == "ENABLE_DEBUG" {
		fmt.Fprintln(os.Stderr, "Enabling debug")
		debugFlag = true
	}
}

func log(format string, args ...interface{}) {
	if !debugFlag {
		return
	}

	fmt.Fprintf(os.Stderr, format, args...)
}
