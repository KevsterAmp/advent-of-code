file = "input.txt"
with open(file, "r") as f:
    lines = f.read().strip().split("\n")
moves = [(i[0], int(i[1:])) for i in lines]
print(moves)


def fix_dial(num):
    while num < 0 or num > 99:
        if num < 0:
            num += 100
        elif num > 99:
            num -= 100
    return num


password = 0
current_num = 50
for direction, move in moves:
    # print(f"{direction=}, {move=}")
    if direction == "L":
        current_num = fix_dial(current_num - move)
    elif direction == "R":
        current_num = fix_dial(current_num + move)
    else:
        raise Exception("invalid direction")
    if current_num == 0:
        password += 1
    # print(f"{current_num=}, {password=}")
    # break
print(f"{current_num=}, {password=}")
