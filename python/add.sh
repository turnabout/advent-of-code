#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <year_last_two_digits> <day>"
    echo ""
    echo "Example:"
    echo "./add.sh 23 8"
    exit 1
fi

year=$1
day=$2

# Create Python file
python_file="y${year}_${day}.py"
cat <<EOF >"$python_file"
use_real_input = False
use_part = 1


def part1(input):
    print(f"Part 1 called with input:\n{input}")


def part2(input):
    print(f"Part 2 called with input:\n{input}")


EOF

echo "Python file created: $python_file"

# Create test input file in the 'cache' directory
cache_dir="cache"
test_input_file="${cache_dir}/${year}_${day}_test.txt"
mkdir -p "$cache_dir"
touch "$test_input_file"

echo "Test input file created: $test_input_file"
