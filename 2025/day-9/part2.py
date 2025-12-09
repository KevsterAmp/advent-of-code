file_name = "test_input.txt"

with open(file_name, "r") as f:
    lines = [
        (int(x), int(y)) for x, y in [x.split(",") for x in f.read().split("\n")[:-1]]
    ]

orig_lines = list(lines)
x_set = {x[0] for x in lines}
y_set = {x[1] for x in lines}
max_x = max(x_set)
max_y = max(y_set)
print(lines)
print(f"{max_x=}, {max_y=}")
print(f"{x_set=}\n{y_set=}")


def get_rec_area(corner1, corner2):
    x1, y1 = corner1
    x2, y2 = corner2

    return (abs(x1 - x2) + 1) * (abs(y1 - y2) + 1)


def is_corner_in_polygon(corner, lines):
    if corner in lines:
        return True
    return False


def get_corners(corner1, corner2):
    x1, y1 = corner1
    x2, y2 = corner2
    corner3 = x1, y2
    corner4 = x2, y1

    if is_in_polygon(*corner3) and is_in_polygon(*corner4):
        return get_rec_area(corner1, corner2)
    return 0


def is_in_polygon(x, y):
    # print(f"{x=}, {y=}")
    flag = False
    # up
    for i in range(y, 0, -1):
        # print(f"up i {i}")
        if (x, i) in orig_lines:
            flag = True
            break
        elif i in y_set:
            # print("up ping")
            limits = get_line(i, "y")
            # print(f"up limits {limits}")
            for limit in limits:
                if x > min(limit) and x < max(limit):
                    flag = True
                    break
        else:
            flag = False
    # print(f"{flag=}")
    if flag is False:
        return False
    # print("up done")

    # down
    for i in range(y, max_y + 1):
        if (x, i) in orig_lines:
            flag = True
            break
        elif i in y_set:
            limits = get_line(i, "y")
            for limit in limits:
                if x > min(limit) and x < max(limit):
                    flag = True
                    break
        else:
            flag = False
    if flag is False:
        return False
    print("down done")

    # left
    for i in range(x, 0, -1):
        if (i, y) in orig_lines:
            flag = True
            break
        elif i in x_set:
            # print(f"{i=}")
            limits = get_line(i, "x")
            for limit in limits:
                if y > min(limit) and y < max(limit):
                    flag = True
                    break
        else:
            flag = False
    if flag is False:
        return False
    print("left done")

    for i in range(x, max_x + 1):
        # print(f"right i {i}")
        if (i, y) in orig_lines:
            flag = True
            break
        elif i in x_set:
            limits = get_line(i, "x")
            for limit in limits:
                if y > min(limit) and y < max(limit):
                    flag = True
                    break
        else:
            flag = False
    # print(f"right flag {flag}")

    if flag:
        return True
    return False


def get_line(n, type):
    print(f"{orig_lines=}")
    pair = []
    if type == "y":
        pair = [x[0] for x in filter(lambda a: a[1] == n, orig_lines)]
    elif type == "x":
        pair = [x[1] for x in filter(lambda a: a[0] == n, orig_lines)]
    if None in pair or len(pair) % 2 == 1:
        print(f"{pair=}")
        raise Exception("there is an issue on the pairs")
    return pair_up(sorted(pair))


def pair_up(pairs):
    output = []
    if len(pairs) > 2:
        for i in range(len(pairs), 2):
            output.append((pairs[i], pairs[i + 1]))
    else:
        output.append((pairs[0], pairs[1]))
    return output


print(f"{len(x_set)=}, {len(y_set)=}, {len(lines)=}")

output = []
output = ((("." * 12) + "\n") * 12).split("\n")[:-1]
output = [list(x) for x in output]


# print(is_in_polygon(*lines[0]))
# for line in output:
#     for char in line:
#         print(char, end="")
#     print("\n")

areas = []
while lines:
    curr = lines.pop(0)
    print(f"{curr=}")
    for line in lines:
        # get rectangle
        areas.append(get_corners(curr, line))
print(f"{max(areas)=}")
print(areas)
