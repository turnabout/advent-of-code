# Advent of Code

Python runner & solutions for Advent of Code.

The runner automatically fetches the "real" input from the Advent of Code website.

## Setup
To fetch input from the website, a `SESSION` environment variable must be set to the user's session value. This can be recovered by logging into [the Advent of Code website](https://adventofcode.com/) and copying the value for the `session` cookie.

## Adding solution files
To define a solution for a given day, run the included `add.sh` script:

```bash
./add.sh <year_last_two_digits> <day>

# Example - add files for 2023 day 8
./add.sh 23 8
```

This will create two files:
* That day's solution file
    * `y<year>_<day>.py`
* Sample input file
    * `input/<year>_<day>_sample.txt`

You must then manually copy, into the second file, the sample input given by the website for that particular day.

The file containing the real input (`input/<year>_<day>_real.txt`) will automatically be created when launching the solution.

## Run
Edit `main.py` to launch the desired solution. This example launches day 7 of 2023:

```python
if __name__ == "__main__":
    run(23, 7)
```

The day's solution file sets two variables to configure its received input, and which function gets invoked:

```python
# Set to `True` to have the invoked solution function receive the real input for that day.
# Else, it will receive the sample input.
use_real_input = False

# Set to `1` or `2` to control whether `part1` or `part2` function get invoked.
used_part = 1
```
