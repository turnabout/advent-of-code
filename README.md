# Advent of Code
Advent of Code solutions. Includes a runner that automatically fetches input from the Advent of Code website.

To define a solution for a given day, create a new file (standard format: `y<year last two digits>_<day>.go` - example: `y23_1.go`), and inside of it, add two new methods to the `Runner` structure named like `Y23_1_1` and `Y23_1_2` (the last digit representing the "part" of the day it solves).

To fetch input from the website, a `SESSION` environment variable must be set to the user's session value. This can be recovered by logging in to [the Advance of Code website](https://adventofcode.com/), copying the value for the `session` cookie.

Example day methods for day 1 (when running `invokeRunnerFunction`, will get invoked & automatically receive their corresponding raw input as a string argument):

```go
package main

func (Runner) Y23_1_1(input string) {
}

func (Runner) Y23_1_2(input string) {
}
```

> Fetched input is kept in files in the `cache` directory to avoid sending unnecessary requests. Delete a file associated with a specific day to force re-fetching from the website.
