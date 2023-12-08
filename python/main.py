import importlib
from get_day_input import get_day_inputs


def import_day_fn(year, day, part):
    module_name = 'y{}_{}'.format(year, day)
    module = importlib.import_module(module_name)

    return getattr(module, "part" + str(part))


def run(year, day, part):
    answer = import_day_fn(year, day, part)(
        *get_day_inputs(year, day)
    )
    print(f"Answer: ", answer)


if __name__ == "__main__":
    run(23, 7, 1)
