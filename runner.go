package main

import (
	"fmt"
	"log"
	"reflect"
)

type Runner struct {
}

// invokeRunnerFunction invokes the runner function for the given
func invokeRunnerFunction(year int, day int) {
	runner := &Runner{}

	// Get method name (Y<short year>_<day>)
	shortYear := 2023 % 2000
	methodName := fmt.Sprintf("Y%d_%d", shortYear, day)

	method := reflect.ValueOf(runner).MethodByName(methodName)

	if method.IsValid() {
		input := getInput(year, day)

		method.Call([]reflect.Value{reflect.ValueOf(input)})
	} else {
		log.Fatalf("Method '%s' was not found", methodName)
	}
}
