# Advent of Code

Python solutions for Advent of Code.

Advent of Code solutions. Includes a runner that automatically fetches input from the Advent of Code website.


To define a solution for a given day, create a new file (format must be: `y<year last two digits>_<day>.py` - example: `y23_1.py`). The file must expose functions named `part1` and `part2`. Reuse this template:

```python
def part1(test_input, real_input):
    use_real_input = False
    input = real_input if use_real_input else test_input


def part2(test_input, real_input):
    use_real_input = False
    input = real_input if use_real_input else test_input
```

To fetch input from the website, a `SESSION` environment variable must be set to the user's session value. This can be recovered by logging in to [the Advance of Code website](https://adventofcode.com/), copying the value for the `session` cookie.

> Fetched input is kept in files in the `cache` directory to avoid sending unnecessary requests. Delete a file associated with a specific day to force re-fetching from the website.

**TODO: Figure a way to get the "test" input - probably will need them added manually**
