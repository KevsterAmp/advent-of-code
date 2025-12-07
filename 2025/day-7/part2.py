import functools

file_name = "input.txt"

with open(file_name, "r") as f:
    lines = [list(x) for x in f.read().split("\n")[:-1]]


# get index of start
idx = lines[0].index("S")
print(f"{idx=}")


hit_count = 0
hits = []
touch = 0
split_beam_runs = 0


print(f"{len(lines)=}")


@functools.lru_cache(maxsize=128)
def split_beam(x, y):
    hits = 0

    for j in range(y + 1, len(lines)):
        if lines[j][x] == "^":
            hits = split_beam(x + 1, j) + split_beam(x - 1, j)
            break

        if j == len(lines) - 1:
            hits += 1
    return hits

    return hits

    # print(f"iter {y}")
    # for line in lines:
    #     print("".join(line))


print(f"{split_beam(idx, 0)=}")
# print(f"{hit_count=}")
# print(f"{touch=}")
