# Advent of Code
Advent of Code solutions written in Go. Includes a small framework for easily fetching input.

## Usage
To automatically fetch input from the Advent of Code website, first log into Advent of Code and find your session cookie. Then, create a file named `session` at the root of the project and paste the contents of the cookie into it.

Then, call the `solutions2021.InvokeSolution` function to invoke whichever solution for the year 2021.

```go
    solutions2021.InvokeSolution(1, 1); // Invokes solution 1 of day 1, 2021
```

In the above example, the function will do some reflection to look up and invoke `Solution2021.S1_1`. It will also look up your problem input on the Advent of Code website to pass to it.

