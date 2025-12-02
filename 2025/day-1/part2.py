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
print(f"{current_num=}, {password=}")
for direction, move in moves:
    print(f"{direction=}, {move=}")
    if direction == "L":
        new_num = fix_dial(current_num - move)

        if current_num == 0:
            password += move // 100
        elif move > current_num:
            password += ((move - current_num) // 100) + 1
        elif current_num == move:
            password += 1
        current_num = new_num

    elif direction == "R":
        password += (current_num + move) // 100
        current_num = fix_dial(current_num + move)
    else:
        raise Exception("invalid direction")
    print(f"{current_num=}, {password=}")
    # break
print(f"{current_num=}, {password=}")
