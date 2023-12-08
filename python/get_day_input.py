from os import environ, path
from urllib.request import Request, urlopen


def get_day_inputs(year, day):
    return get_test_day_input(year, day).strip(), get_cached_real_day_input(year, day).strip()


def get_test_day_input(year, day):
    file_path = f"cache/{year}_{day}_test.txt"
    if not path.exists(file_path):
        raise ValueError(f"test day input file at '{file_path}' does not exist")

    with open(file_path, "r") as file:
        return file.read()


def get_cached_real_day_input(year, day):
    cache_file_path = f"cache/{year}_{day}.txt"

    # Attempt to read from cache
    if path.exists(cache_file_path):
        with open(cache_file_path, "r") as file:
            return file.read()

    # Get from website
    input = get_real_day_input(year, day)

    # Store result in cache
    with open(cache_file_path, "w") as file:
        file.write(input)

    return input


def get_real_day_input(year, day):
    session = environ.get("SESSION", "")

    if session == "":
        raise ValueError("Must set the 'SESSION' environment variable to get a day's input.")

    # Get content from website
    request = Request(
        f"https://adventofcode.com/20{year}/day/{day}/input",
        None,
        {
            "Cookie": "session={}".format(session)
        }
    )
    with urlopen(request) as response:
        content = response.read().decode("utf-8").strip()

    return content
