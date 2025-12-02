file_name = "input.txt"
with open(file_name, "r") as f:
    line = f.read().strip("\n")

product_ranges = line.split(",")

sum = 0
for product_range in product_ranges:
    x, y = [int(k) for k in product_range.split("-")]
    print(len(str(y)))
    for i in range(x, y + 1):
        str_i = str(i)
        digit_len = len(str_i)
        if len(str(i)) % 2 == 1:
            continue
        middle_point = digit_len // 2
        if str_i[:middle_point] == str_i[middle_point:]:
            print(f"{str_i} is fake")
            sum += i

print(f"{sum=}")
