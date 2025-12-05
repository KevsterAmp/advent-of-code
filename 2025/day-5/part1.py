file_name = "input.txt"
with open(file_name, "r") as f:
    fresh, ids = [x.split("\n") for x in f.read().strip().split("\n\n")]
# print(fresh)
# print(ids)


# fresh_range = []
fresh_ids_count = 0
fresh_ids = []
for x in fresh:
    start, end = [int(j) for j in x.split("-")]
    for id in ids:
        if id in fresh_ids:
            continue
        elif int(id) in range(start, end + 1):
            fresh_ids_count += 1
            fresh_ids.append(id)

print(f"{fresh_ids_count=}")
