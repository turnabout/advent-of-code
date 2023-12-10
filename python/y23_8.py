import math
from functools import reduce

use_real_input = True
used_part = 2


def parse(input, pt2=False):
    lines = input.split("\n")

    # Get commands and graph
    commands = list(lines[0])
    current_node, graph = parse_graph(lines[2:], pt2)
    return commands, current_node, graph


def parse_graph(lines: list, pt2=False):
    graph = {}
    starting_node = None
    starting_nodes = []

    for line in lines:
        node_label, right_str = line.split(" = ")
        left_right_str = right_str[1:-1]
        node_left, node_right = left_right_str.split(", ")
        graph[node_label] = {
            "L": node_left,
            "R": node_right,
            "label": node_label
        }

        # Part 1: "AAA" is starting point
        if node_label == "AAA" and not pt2:
            starting_node = graph[node_label]

        # Part 2: All nodes ending with "A" are starting points
        if node_label[2] == "A" and pt2:
            starting_nodes.append(graph[node_label])

    ret_val = starting_node if pt2 is False else starting_nodes

    return ret_val, graph


def get_next_node(current_node, command, graph):
    current_node_label = current_node["label"]
    next_node_label = current_node[command]
    return graph[next_node_label]


def part1(input):
    commands, current_node, graph = parse(input)

    # Keep track of steps taken
    steps_taken = 0

    i = 0
    while True:
        # Start commands from the beginning
        if i >= len(commands):
            i = 0

        # If current node is ZZZ, done
        if current_node["label"] == "ZZZ":
            break

        current_node = get_next_node(
            current_node,
            commands[i],
            graph
        )

        steps_taken += 1
        i += 1

    return steps_taken


def part2(input):
    commands, current_nodes, graph = parse(input, pt2=True)

    steps_taken_list = [0 for _ in current_nodes]
    done_list = [False for _ in current_nodes]

    steps_taken = 0
    i = 0

    while True:
        # Get current command
        if i >= len(commands):
            i = 0
        command = commands[i]
        i += 1

        # All nodes are done (reached the end at least once)
        if all(done for done in done_list):
            break

        # Update all current nodes
        for j, current_node in enumerate(current_nodes):
            current_nodes[j] = get_next_node(current_node, command, graph)
            if not done_list[j]:
                steps_taken_list[j] += 1

            if current_nodes[j]["label"][2] == "Z":
                done_list[j] = True

        steps_taken += 1

    print("Steps taken:")
    print(steps_taken_list)

    # Calculate LCM between the length of all the paths to get the answer
    # (which will be the time when all the current nodes end up on an end node at once)
    return lcm_of_sequence(steps_taken_list)


def lcm_of_sequence(sequence: list):
    return reduce(math.lcm, sequence)