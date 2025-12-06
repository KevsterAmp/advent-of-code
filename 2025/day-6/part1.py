file_name = "input.txt"
with open(file_name, "r") as f:
    lines = [x.strip().split() for x in f.read().strip().split("\n")]

overall = 0
for i in range(len(lines[0])):
    op = lines[-1][i]
    total = 1 if op == "*" else 0
    for j in range(len(lines) - 1):
        if op == "*":
            print(f"{int(lines[j][i])=}")
            total *= int(lines[j][i])
        else:
            total += int(lines[j][i])
    overall += total

    print(f"{total=}")
print(f"{overall=}")
