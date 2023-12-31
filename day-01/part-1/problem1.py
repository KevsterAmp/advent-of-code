def main():
    with open("puzzle_input.txt", "r") as f:
        sum = 0
        for line in f:
            num = func_list(line)
            sum += num
    print(sum)

def func_list(line):
    indices = [i for i, element in enumerate(line) if element.isdigit()]
    return int(line[min(indices)] + line[max(indices)]) 



main()