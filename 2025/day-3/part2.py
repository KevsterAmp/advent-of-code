file = "input.txt"
with open(file, "r") as f:
    lines = f.read().strip().split("\n")
lines

sum_joltage = 0
for bank in lines:
    # print(bank)
    bank_list = list(bank)
    max_len = 12
    str_int = ""
    for j in range(12):
        curr_highest = 0
        idx = 0
        # print(f"{bank_list=}")
        # print(f"{(len(bank_list) - max_len)=}")
        for i in range(len(bank_list) - max_len + 1):

            # print(bank_list[i], end="")
            if int(bank_list[i]) > curr_highest:
                curr_highest = int(bank_list[i])
                idx = i
        # print("\n")
        bank_list = bank_list[slice(idx + 1, len(bank_list))]
        max_len -= 1
        # print(f"{idx=}")
        # print(f"{curr_highest=}")
        str_int += str(curr_highest)
    print(f"{str_int=}")
    sum_joltage += int(str_int)
print(f"{sum_joltage=}")
