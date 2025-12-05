file_name = "input.txt"
with open(file_name, "r") as f:
    fresh, ids = [x.split("\n") for x in f.read().strip().split("\n\n")]

fresh_ranges = sorted(
    [[int(start), int(end)] for start, end in [x.split("-") for x in fresh]]
)


total_ranges = []
total_count = 0
for i, (start, end) in enumerate(fresh_ranges):
    while True:
        if i == len(fresh_ranges) - 1:
            break
        next_start, next_end = fresh_ranges[i + 1]
        if next_start > end:
            break
        else:
            end = max(end, next_end)
            fresh_ranges.pop(i)
    total_ranges.append((start, end))
    total_count += end - start + 1

print(f"{total_count=}")
