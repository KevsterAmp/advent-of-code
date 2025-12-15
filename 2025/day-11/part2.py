import functools

file_name = "input.txt"

with open(file_name, "r") as f:
    lines = {
        x: y.split(" ")
        for x, y in list(map(lambda x: x.split(": "), f.read().split("\n")[:-1]))
    }

# print(lines)

start = lines.get("you")
# print(f"{start=}")

cache = {}


@functools.cache
def travel(key, dac, fft):
    global lines
    count = 0
    for out in lines.get(key, []):
        curr_dac = dac
        curr_fft = fft
        # print(f"{key=} {out=} {curr_dac=} {curr_fft=}")

        if "dac" == out:
            curr_dac = 1
        if "fft" == out:
            curr_fft = 1
        if "out" in lines.get(out, []):
            if dac == 1 and fft == 1:
                count += 1
            continue
        count += travel(out, curr_dac, curr_fft)

    # print(f"{count=}")
    return count


out = travel("svr", 0, 0)
print(out)
