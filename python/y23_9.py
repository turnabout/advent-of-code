use_real_input = True
used_part = 2


# 1. Create new sequences from the difference of each step, until they are all 0
# - Add 0 to the end of the list of zeroes
#   - There is now a placeholder in every sequences above it
# 2. Fill in placeholders from the bottom up:
#   - Increase the value below it by the value left of it
# 3. The new last value is the predicted next value in the history
# 4. Add up all the predicted next values together


def create_sequences_list(starting_values: list):
    out_lists = [starting_values]

    create_sequence_helper(starting_values, out_lists)

    return out_lists


def create_sequence_helper(values: list, out_list: list):
    new_sequence = create_sequence_from_values(values)

    out_list.append(new_sequence)

    # If all values in new sequence are NOT 0, keep going
    if not all(val == 0 for val in new_sequence):
        create_sequence_helper(new_sequence, out_list)


def create_sequence_from_values(values: list):
    if len(values) == 1:
        return [0]

    new_sequence = []
    for i, current_value in enumerate(values):
        # Get next value
        # (Reached end, no more "next value")
        if i + 1 >= len(values):
            break
        next_value = values[i + 1]

        # Record new value (difference between this & next value)
        new_sequence.append(
            next_value - current_value
        )

    return new_sequence


def parse_lines(input: str):
    # String to input split by lines
    lines = input.split("\n")

    # Transform each line into a list of ints
    return [
        [int(num_str) for num_str in line.split(" ")]
        for line in lines
    ]


def get_predicted_value(sequences: list, after=True):
    sequences.reverse()
    carry = 0

    for sequence in sequences:
        if after:
            carry += sequence[-1]
        else:
            carry = sequence[0] - carry

    return carry


def part1(input):
    lines = parse_lines(input)
    sequences_list = [create_sequences_list(line) for line in lines]

    answer = 0
    for sequences in sequences_list:
        answer += get_predicted_value(sequences)

    return answer


def part2(input):
    lines = parse_lines(input)
    sequences_list = [create_sequences_list(line) for line in lines]

    answer = 0
    for sequences in sequences_list:
        answer += get_predicted_value(sequences, False)

    return answer


