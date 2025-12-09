file_name = "test_input.txt"

with open(file_name, "r") as f:
    lines = [
        (int(x), int(y)) for x, y in [x.split(",") for x in f.read().split("\n")[:-1]]
    ]

print(lines)


def get_rec_area(corner1, corner2):
    x1, y1 = corner1
    x2, y2 = corner2

    return (abs(x1 - x2) + 1) * (abs(y1 - y2) + 1)


areas = []
while lines:
    curr = lines.pop(0)
    print(f"{curr=}")
    for line in lines:
        # get rectangle
        areas.append(get_rec_area(curr, line))

print(max(areas))
print(areas)

# print(get_rec_area((11, 7), (2, 3)))
