import re


def is_fake(str_i, j):
    segments = re.findall(f".{{{j}}}", str_i)
    if len(segments) > 1 and len(set(segments)) == 1 and i not in fakes:
        fakes.append(i)
        print(segments)
        return True
    return False


file_name = "input.txt"
with open(file_name, "r") as f:
    line = f.read().strip("\n")

product_ranges = line.split(",")

sum = 0
fakes = []
for product_range in product_ranges:
    x, y = [int(k) for k in product_range.strip().split("-")]
    print(x, y)
    for i in range(x, y + 1):
        str_i = str(i)
        digit_len = len(str_i)
        if len(str(i)) == 9:
            for a in [1, 3]:
                if is_fake(str_i, a):
                    sum += i
                    print(f"{str_i} is fake")

        elif len(str(i)) % 2 == 1:
            if is_fake(str_i, 1):
                sum += i
                print(f"{str_i} is fake")

        else:
            middle_point = digit_len // 2
            # if str_i[:middle_point] == str_i[middle_point:]:
            for j in [1, 2, middle_point]:
                if is_fake(str_i, j):
                    sum += i
                    print(f"{str_i} is faker")
                    # print(segments)
                    # print(j)
                    # print(f"{digit_len=}")

            # print(segments)
print(f"{sum=}")
