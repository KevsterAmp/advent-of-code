import re


with open("test.in", "r") as file:
    lines = file.read().split("\n")

    
def part1():
    direction = lines[0]

    def calibrate(lines): 
        guide = lines[2:]
        orig, l, r = [], [], []
        pattern = r'\(([^,]+), ([^)]+)\)'
        for line in guide:
            # get orig, l, r
            temp = line.split("=")
            orig.append(temp[0].strip())
            x, y = temp[1].replace("(", "").replace(")", "").split(",")
            l.append(x.strip())
            r.append(y.strip())

        return orig, l, r

    def get_steps(orig, l, r):
        steps = 0
        curr = orig[0]
        print(f"end: {orig[-1]}")
        print(f"Curr: {curr}")
        while curr != orig[-1]:
            # get the index of the orig placement depending on the guide
            for x in direction:
                # for debugging

                if x == "L":
                    lr = l
                elif x == "R":
                    lr = r
                else: 
                    return "error"
                # get the index of curr
                i = orig.index(curr)
                # using that index, get the direction depending if L or R
                points_to = lr[i]
                curr_i = orig.index(points_to)
                curr = orig[curr_i]
                steps += 1
        
        return steps

    orig, l, r = calibrate(lines)
    return get_steps(orig, l, r)


def part2():
    pass

if __name__ in "__main__":
    print(f"Part 1: {part1()}")
    print(f"Part 2: {part2()}")