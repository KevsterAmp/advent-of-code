from itertools import cycle
import math
import re


with open("input.in", "r") as file:
    lines = file.read()
    direction, maps = lines.split("\n\n")
    direction = cycle(0 if x == "L" else 1 for x in direction)
    clean_maps = {}
    regex = r'(\w{3}) = \((\w{3}), (\w{3})\)'

    for start, left, right in re.findall(regex, maps):
        clean_maps[start] = [left, right]

    
def part1():
    output = 0
    node = "AAA"
    for x in direction:
        output += 1
        node = clean_maps[node][x]
        if node == "ZZZ":
            break

    return output # type: ignore

def part2():
    node_a = [node for node in list(clean_maps.keys()) if node[-1] == "A"] # the starting nodes
    steps4node = []
    for node in node_a:
        steps = 0
        for x in direction:
            steps += 1
            node = clean_maps[node][x]

            if node[-1] == "Z":
                steps4node.append(steps)
                break

    return math.lcm(*steps4node)


if __name__ in "__main__":
    print(f"Part 1: {part1()}")
    print(f"Part 2: {part2()}")
