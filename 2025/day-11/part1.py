file_name = "input.txt"

with open(file_name, "r") as f:
    lines = {
        x: y.split(" ")
        for x, y in list(map(lambda x: x.split(": "), f.read().split("\n")[:-1]))
    }

print(lines)

start = lines.get("you")
print(f"{start=}")


def travel(key, lines):
    count = 0
    # if lines.get(key) is None:
    #     raise Exception(f"{key} not found.")
    for out in lines.get(key):
        # print(f"{key=} {out=}")
        if "out" in lines.get(out):
            # print("out!")
            count += 1
        else:
            count += travel(out, lines)

    return count


out = travel("svr", lines)
print(out)
