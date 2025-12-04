file_name = "input.txt"
with open(file_name, "r") as f:
    lines = f.read().strip().split("\n")

max_x = len(lines[0])
max_y = len(lines)


def check_adj(x, y):
    adj = 0
    xs = [x, x + 1, x - 1]
    ys = [y, y + 1, y - 1]

    for i in xs:
        for j in ys:
            if i == x and j == y:
                continue
            elif i < 0 or i > max_x - 1:
                continue
            elif j < 0 or j > max_y - 1:
                continue
            else:
                if lines[j][i] == "@":
                    adj += 1
    return adj

    # [((y - 1, y + 1), (x, x + 1, x - 1))]


def remove_rolls(rolls_loc):
    for x, y in rolls_loc:
        lines[y][x] = "."


curr_rolls = 0
rolls_loc = []
for y, line in enumerate(lines):
    for x, c in enumerate(line):
        if c == "@":
            if check_adj(x, y) < 4:
                curr_rolls += 1
print(f"{curr_rolls=}")
