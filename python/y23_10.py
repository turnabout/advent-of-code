use_real_input = True
used_part = 2


# ANSI escape codes for text color
class TextColors:
    RESET = "\033[0m"
    RED = "\033[91m"
    GREEN = "\033[92m"
    YELLOW = "\033[93m"
    BLUE = "\033[94m"
    MAGENTA = "\033[95m"
    CYAN = "\033[96m"


def print_color(text: str, color: str):
    print(f"{color}{text}{TextColors.RESET}", end="")


actual_s_char = "7"
s_char_coordinates = ""


"""
Keeps track of all the connections available for each type of tile.
"""
tile_connections = {
    ".": [],
    "|": ["t", "b"],
    "-": ["l", "r"],
    "F": ["b", "r"],
    "7": ["l", "b"],
    "L": ["t", "r"],
    "J": ["l", "t"],

    "S": []
}


pretty_print_chars = {
    "|": "│",
    "-": "─",
    "F": "┌",
    "7": "┐",
    "L": "└",
    "J": "┘",
}

def parse_graph(input: str):
    """
    Parses the raw input string into a graph (list of lists).
    The "innermost" lists represent rows.
    """
    return [
        [char for char in line]
        for line in input.split("\n")
    ]


def get_starting_tile(graph: list):
    """
    Finds the starting tile ("S") in the graph.
    """
    for y, line in enumerate(graph):
        if "S" in line:
            x = line.index("S")
            graph[y][x] = actual_s_char
            s_char_coordinates = f"{x}_{y}"
            return get_tile(graph, x, y)

    return None


def get_tile(graph: list, x: int, y: int):
    """
    Gets the value of the tile at the given coordinates.
    """
    if y < 0 or y >= len(graph) or x < 0 or x >= len(graph[0]):
        return None
    return {
        "tile": graph[y][x],
        "x": x,
        "y": y
    }


def get_surrounding_tiles_list(graph: list, middle_tile):
    surrounding_tiles_obj = get_surrounding_tiles(graph, middle_tile)
    return [
        surrounding_tiles_obj["top"],
        surrounding_tiles_obj["bottom"],
        surrounding_tiles_obj["left"],
        surrounding_tiles_obj["right"]
    ]


def get_surrounding_tiles(graph: list, middle_tile):
    x = middle_tile["x"]
    y = middle_tile["y"]

    """
    Gets the value of the tiles surrounding the given coordinates.
    """
    return {
        "top": get_tile(graph, x, y - 1),
        "right": get_tile(graph, x + 1, y),
        "bottom": get_tile(graph, x, y + 1),
        "left": get_tile(graph, x - 1, y),
    }


def get_connected_surrounding_tiles(graph: list, middle_tile):
    surrounding_tiles = get_surrounding_tiles(graph, middle_tile)

    res = []
    if surrounding_tiles["right"] is not None and horizontal_tiles_connect(middle_tile["tile"], surrounding_tiles["right"]["tile"]):
        res.append(surrounding_tiles["right"])

    if surrounding_tiles["left"] is not None and horizontal_tiles_connect(surrounding_tiles["left"]["tile"], middle_tile["tile"]):
        res.append(surrounding_tiles["left"])

    if surrounding_tiles["top"] is not None and vertical_tiles_connect(surrounding_tiles["top"]["tile"], middle_tile["tile"]):
        res.append(surrounding_tiles["top"])

    if surrounding_tiles["bottom"] is not None and vertical_tiles_connect(middle_tile["tile"], surrounding_tiles["bottom"]["tile"]):
        res.append(surrounding_tiles["bottom"])

    return res


def horizontal_tiles_connect(left: str, right: str):
    return "r" in tile_connections[left] and "l" in tile_connections[right]


def vertical_tiles_connect(top: str, bottom: str):
    return "b" in tile_connections[top] and "t" in tile_connections[bottom]


def adjust_starting_tile_connections(directions: list):
    for direction in directions:
        tile_connections["S"].append(direction)


def bfs(graph, starting_tile):
    visited = []
    queue = [] # Queue contains elements looking like "<x>_<y>" - ex: "10_4"

    visited.append(f"{starting_tile['x']}_{starting_tile['y']}")
    queue.append(starting_tile)
    steps = 0

    while len(queue) > 0:
        current_tile = queue.pop(0)

        # Get neighbors
        for neighbor in get_connected_surrounding_tiles(graph, current_tile):
            neighbor_visited_key = f"{neighbor['x']}_{neighbor['y']}"
            if neighbor_visited_key in visited:
                continue

            # Enqueue neighbor
            visited.append(neighbor_visited_key)
            queue.append(neighbor)
            steps += 1

    print(f"Total steps: {steps}")

    return visited


def bfs_outer(graph, starting_tile, excluded_tiles):
    visited = []
    queue = []

    visited.append(f"{starting_tile['x']}_{starting_tile['y']}")
    queue.append(starting_tile)

    while len(queue) > 0:
        current_tile = queue.pop(0)

        # Get neighbors
        for neighbor in get_surrounding_tiles_list(graph, current_tile):
            if neighbor is None:
                continue

            neighbor_visited_key = f"{neighbor['x']}_{neighbor['y']}"
            if neighbor_visited_key in visited or neighbor_visited_key in excluded_tiles:
                continue

            # Enqueue neighbor
            visited.append(neighbor_visited_key)
            queue.append(neighbor)

    return visited


def part1(input: str):
    # TODO: Adjust depending on input
    if use_real_input:
        actual_s_char = "J"
    else:
        actual_s_char = "F"

    graph = parse_graph(input)
    starting_tile = get_starting_tile(graph)

    visited = bfs(graph, starting_tile)
    return int(((len(visited) - 2) / 2) + 1)

    # print(starting_tile)
    # print(get_connected_surrounding_tiles(graph, starting_tile))
    # print(get_surrounding_tiles(graph, 1, 1))
    # print(get_connected_surrounding_tiles(graph, get_tile(graph, 1, 2)))


def part2(input: str):
    if use_real_input:
        adjust_starting_tile_connections(["t", "l"])
        actual_s_char = "J"
    else:
        actual_s_char = "7"

    graph = parse_graph(input)

    starting_tile = get_starting_tile(graph)
    loop_visited_tiles = bfs(graph, starting_tile)

    visited_outer = bfs_outer(graph, get_tile(graph, 0, 0), loop_visited_tiles)
    a = s_char_coordinates

    # Loop line
    for y, line in enumerate(graph):
        # Loop tiles on line
        for x, tile_str in enumerate(line):
            tile_str_key = f"{x}_{y}"

            if tile_str_key == s_char_coordinates:
                print_color("#", TextColors.GREEN)
            if tile_str_key in loop_visited_tiles:
                print_color(pretty_print_chars[tile_str], TextColors.CYAN)
            elif tile_str_key in visited_outer:
                print_color(".", TextColors.RED)
            else:
                print(".", end="")
        print("")

