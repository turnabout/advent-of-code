package solutions2021

import (
	"fmt"
	"github.com/turnabout/advent-of-code-2021/pkg/input"
	"reflect"
)

type Solution struct{}

func InvokeSolution(day int, part int) {
	s := Solution{}
	method := reflect.ValueOf(s).MethodByName(
		fmt.Sprintf("S%d_%d", day, part),
	).Interface().(func(string) string)

	fmt.Println(method(input.Fetch2021Input(day)))
}
