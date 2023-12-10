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
used_part = 1


def part1(input):
    print(f"Part 1 called with input:\n{input}")


def part2(input):
    print(f"Part 2 called with input:\n{input}")


EOF

echo "Python file created: $python_file"

# Create sample input file in the 'cache' directory
cache_dir="cache"
sample_input_file="${cache_dir}/${year}_${day}_sample.txt"
mkdir -p "$cache_dir"
touch "$sample_input_file"

echo "Sample input file created: $sample_input_file"
