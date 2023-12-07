def parse(input):
    """
    Returns zipped lists of corresponding times and distances.
    """
    lines = input.split("\n")

    times = list(
        map(
            int,
            lines[0].split()[1:]
        )
    )

    distances = list(
        map(
            int,
            lines[1].split()[1:]
        )
    )

    return list(zip(times, distances))


def parse_second_part(input):
    separated = parse(input)
    transposed = zip(*separated)
    totalTimeInts, recordInts = map(tuple, transposed)

    totalTime = int(''.join(map(str, totalTimeInts)))
    record = int(''.join(map(str, recordInts)))

    return totalTime, record


def calc_distance(total_time, held_time):
    mm_per_ms = held_time
    ms_left = total_time - held_time

    # No distance if no time left
    if ms_left <= 0:
        return 0

    return ms_left * mm_per_ms


def part1(test_input, real_input):
    use_real_input = True
    input = real_input if use_real_input else test_input

    parsed = parse(input)
    answer = 1

    for item in parsed:
        record = item[1]
        held_time = 1
        possible_ways = 0

        while True:
            distance = calc_distance(item[0], held_time)
            if distance == 0:
                break
            held_time += 1
            if distance > record:
                possible_ways += 1

        print(f'possible ways on {item}: {possible_ways}')
        answer *= possible_ways

    return answer


def find_first_non_winning_held_time(record, total_time, initial_held_time, held_time_update):
    held_time = initial_held_time

    # Find a first "valid" value
    while True:
        distance = calc_distance(total_time, held_time)

        # Found no winning distance
        if distance == 0:
            print(f'look_for_first_non_winning_time: got a "0" distance for held_time {held_time}')

        # Found a non-winning distance
        if distance <= record or distance == 0:
            return held_time

        # Unfortunately won - take a leap
        held_time += held_time_update


def find_first_winning_held_time(record, total_time, initial_held_time, held_time_update):
    held_time = initial_held_time

    # Find a first "valid" value
    while True:
        distance = calc_distance(total_time, held_time)

        # Found a winning distance!
        if distance > record:
            return held_time

        # Didn't win - take a leap
        held_time += held_time_update


def part2(test_input, real_input):
    use_real_input = False

    input = real_input if use_real_input else test_input

    # Get amount we should update held time when searching, & get total time / record
    held_time_update = 100000 if use_real_input else 100
    total_time, record = parse_second_part(input)

    # Find a starting point for first winning held time
    first_winning_held_time = find_first_winning_held_time(record, total_time, 1, held_time_update)

    # Find start for "winning range beginning" search
    beginning_search_start = find_first_non_winning_held_time(
        record,
        total_time,
        first_winning_held_time,
        -held_time_update
    )

    # Find start for "winning range ending" search
    ending_search_start = find_first_non_winning_held_time(
        record,
        total_time,
        first_winning_held_time,
        held_time_update
    )

    # Find actual start of range
    winning_range_start = find_first_winning_held_time(record, total_time, beginning_search_start, 1)
    winning_range_end = find_first_winning_held_time(record, total_time, ending_search_start, -1)

    answer = winning_range_end - winning_range_start + 1
    print(f'there are {answer} ways to win')

    return answer

