file_name = "input.txt"

with open(file_name, "r") as f:
    lines = [list(x) for x in f.read().split("\n")[:-1]]


# get index of start
idx = lines[0].index("S")
print(f"{idx=}")


hit_count = 0
hits = []


def split_beam(x, y):
    global hits, hit_count
    for j in range(y + 1, len(lines) - 1):
        if lines[j][x] == "^":
            lines[j][x] = "x"
            hit_count += 1
            split_beam(x + 1, j)
            split_beam(x - 1, j)
            break
        elif lines[j][x] == "x":
            break
        else:
            lines[j][x] = "|"
    print(f"iter {y}")
    for line in lines:
        print("".join(line))


split_beam(idx, 0)
print(f"{hit_count=}")
