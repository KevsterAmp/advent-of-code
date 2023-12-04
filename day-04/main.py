def main():
    output_p1 = 0
    match_list = []
    with open("input", "r") as file:
        lines = file.readlines()
        cards_instances = [1] * len(lines)
        for line in lines:
            # part 1
            matches_num = check_matches_num(line)
            if matches_num > 0:
                output_p1 += pow(2, matches_num -1)

            # part 2
            match_list.append(matches_num)

    # add the copies to each other
    for i, num_match in enumerate(match_list):
        if num_match > 0:
            for _ in range(cards_instances[i]):
                for j in range(1, num_match+1):
                    if (i+j) > len(match_list):
                        break
                    cards_instances[i+j] += 1

    print(f"part 1: {output_p1}")
    print(f"part 2: {sum(cards_instances)}")


def check_matches_num(line):
    numbers = line.split(":")[1].split("|")
    winning, curr = map(str.split, numbers)
    return len([x for x in curr if x in winning])


main()
