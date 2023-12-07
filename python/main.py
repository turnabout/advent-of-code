import importlib


def import_day_fn(year, day, part):
    module_name = 'y{}_{}'.format(year, day)
    module = importlib.import_module(module_name)

    return getattr(module, "part" + str(part))


def get_day_inputs(year, day):
    test_input = """Time:      7  15   30
    Distance:  9  40  200"""

    real_input = """Time:        51     69     98     78
    Distance:   377   1171   1224   1505"""

    return test_input, real_input


def run(year, day, part):
    answer = import_day_fn(year, day, part)(
        *get_day_inputs(year, day)
    )
    print(f'Answer: ', answer)


if __name__ == "__main__":
    run(23, 6, 1)
