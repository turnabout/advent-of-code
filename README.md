# Advent of Code
Advent of Code solutions. Includes a runner that automatically fetches input from the Advent of Code website.

To define a solution for a given day, create a new file (standard format: `y<year last two digits>_<day>.go` - example: `y23_1.go`), and inside of it, add a new method to the `Runner` structure named like `Y23_1`.

To fetch input from the website, a `SESSION` environment variable must be set to the user's session value. This can be recovered by logging in to [the Advance of Code website](https://adventofcode.com/), copying the value for the `session` cookie.

Example day method (will automatically receive the raw input as a string argument):

```go
package main

import "fmt"

func (Runner) Y23_1(input string) {
    fmt.Println("day 1")
}
```
