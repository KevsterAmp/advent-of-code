
def main():
    with open("puzzle_input.txt", "r") as f:
        sum = 0
        for line in f:
            num = func_list(line)
            sum += num
    print(sum)


def some_random_func(line: str) -> int:
    indices = {i: element for i, element in enumerate(line) if element.isdigit()}
    return int(indices[min(indices)] + indices[max(indices)]) 


def func_list(line):
    indices = [i for i, element in enumerate(line) if element.isdigit()]
    return int(line[min(indices)] + line[max(indices)]) 



main()