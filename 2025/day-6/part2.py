import re

file_name = "input.txt"
with open(file_name, "r") as f:
    lines = f.read()
print(lines)
print(list(lines))
print(lines.index("\n"))
print(len(lines))

# get max len
len_lines = []
print(lines.split("\n"))
for line in lines.split("\n")[:-1]:
    len_lines.append(len(line))
max_len = max(len_lines)
print(max_len)

# add up spaces on the rightside if len is lower than max
normalized = []
for line in lines.split("\n")[:-1]:
    if len(line) <= max_len:
        diff = max_len - len(line)
        normalized.append(line + (" " * diff))
    else:
        raise Exception("There's an issue")

print(f"{normalized=}")

# get length of problems for normalization
match = re.findall(r"[*+] +", normalized[-1])
print(match)
match_len = [len(x) for x in match]
print(f"{match_len=}")

# use length to get the range per problem
length_len = []
end = None
for i, j in enumerate(match_len):
    start = end + 1 if end else 0

    if i == len(match_len) - 1:
        end = start + j
    else:
        end = start + j - 1
    length_len.append([start, end])
print(f"{length_len=}")

normalized_list = []
for line in normalized:
    norm = []
    for start, end in length_len:
        norm.append(line[start:end])
    # for i in range(0, len(line), 4):
    #     norm.append(line[i : i + 3])
    normalized_list.append(norm)
print(f"{normalized_list=}")
print(f"{normalized_list[0]=}")

# get the vert num and compute the problem
overall = 0
for i in range(len(normalized_list[0])):

    op = normalized_list[-1][i].strip()
    total = 1 if op == "*" else 0
    diff = length_len[i][1] - length_len[i][0]
    for j in range(diff):
        vert_num = ""
        for c in range(len(normalized_list) - 1):
            print(f"{c=} {i=} {j=}")
            vert_num += normalized_list[c][i][j]

        print(f"{vert_num=}")
        if vert_num.strip() == "":
            vert_num = 0

        if op == "*":
            total *= int(vert_num)
        else:
            total += int(vert_num)

    print(f"{total=}")
    overall += total

# print(f"{total_vert_list=}")
print(f"{overall=}")
