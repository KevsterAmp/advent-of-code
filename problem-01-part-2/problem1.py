def main():
    sum = 0
    with open("new.txt", "r") as f:
        for line in f:
            num_1 = func_list(line)
            sum += num_1
    print(sum)


def func_list(line):
    line = (line.replace("one", "o1e").replace("two", "t2o").replace("three", "t3e").replace("four", "f4r").replace("five", "f5e")
    .replace("six", "s6x").replace("seven", "s7n").replace("eight", "e8t").replace("nine", "n9e"))
    indices = [element for element in line if element.isdigit()]
    return int(indices[0] + indices[-1])

main()