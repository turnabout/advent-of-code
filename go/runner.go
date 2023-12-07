package _go

import (
	"fmt"
	"log"
	"reflect"
)

type Runner struct {
}

// invokeRunnerFunction invokes the runner function for the given
func invokeRunnerFunction(year int, day int, part int) {
	runner := &Runner{}

	// Get method name (Y<short year>_<day>_<part>)
	shortYear := 2023 % 2000
	methodName := fmt.Sprintf("Y%d_%d_%d", shortYear, day, part)

	method := reflect.ValueOf(runner).MethodByName(methodName)

	if method.IsValid() {
		input := getInput(year, day)

		method.Call([]reflect.Value{reflect.ValueOf(input)})
	} else {
		log.Fatalf("Runner method '%s' was not found", methodName)
	}
}
