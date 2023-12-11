import importlib
from get_day_inputs import get_day_inputs


def get_solution_fn_and_input(year, day):
    # Import module for the given day
    module_name = 'y{}_{}'.format(year, day)
    module = importlib.import_module(module_name)

    # Get module's configured input / part
    use_real_input = getattr(module, "use_real_input")
    used_part = getattr(module, "used_part")

    # Get fn
    fn = getattr(module, "part" + str(used_part))

    # Get input
    sample_input, real_input = get_day_inputs(year, day)
    input = real_input if use_real_input else sample_input

    return fn, input


def run(year, day):
    fn, input = get_solution_fn_and_input(year, day)
    answer = fn(input)
    print(f"Answer: ", answer)

if __name__ == "__main__":
    run(23, 9)
