file = "input.txt"
with open(file, "r") as f:
    lines = f.read().strip().split("\n")
lines

sum = 0
for bank in lines:
    # print(bank)
    curr_highest = 0
    bank_list = list(bank)
    for i in range(len(bank) - 1):
        bank_list.pop(0)
        for num in bank_list:
            joltage = int(bank[i] + num)
            # print(joltage)
            # print(curr_highest)
            if joltage > curr_highest:
                curr_highest = joltage
                # print(f"new {curr_highest=}")
    # print(f"{curr_highest=}")
    sum += curr_highest
print(sum)
