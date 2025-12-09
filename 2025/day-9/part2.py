from shapely import Polygon, box

file_name = "input.txt"

with open(file_name, "r") as f:
    points = [
        (int(x), int(y)) for x, y in [x.split(",") for x in f.read().split("\n")[:-1]]
    ]

print(points)

dots = list(points)


def get_rec_area(x1, y1, x2, y2):
    return (abs(x1 - x2) + 1) * (abs(y1 - y2) + 1)


rects = []
while dots:
    x1, y1 = dots.pop(0)
    for dot in dots:
        x2, y2 = dot
        min_x = min(x1, x2)
        min_y = min(y1, y2)
        max_x = max(x1, x2)
        max_y = max(y1, y2)
        rects.append([min_x, min_y, max_x, max_y])

areas = [get_rec_area(*coords) for coords in rects]

poly = Polygon(points)

max_area = 0
for area, rect in zip(areas, rects):
    if poly.contains(box(*rect)) and area > max_area:
        max_area = area

print(f"{max_area=}")
