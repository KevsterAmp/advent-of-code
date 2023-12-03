import re


def is_symbol(char):
    # Check if the character is not alphanumeric or a whitespace character
    return not (char.isalpha() or char.isdigit() or char.isspace() or char == ".")


def main():
    part1_sum = 0
    part2_sum = 0
    with open("input", "r") as file:
        lines = calibrate(file.read().split("\n"))
        for i, line in enumerate(lines):
            for j, char in enumerate(line):
                # part 1
                if is_symbol(char):
                    adj_curr_line = [line[j-1], line[j+1]]
                    adj_top_line = list(lines[i-1][j-1:j+2])
                    adj_below_line = list(lines[i+1][j-1:j+2])
                    all_adj = adj_curr_line + adj_top_line + adj_below_line  
                    
                    for x in remove_adjacent_duplicates(all_adj):
                        if x.isdigit():
                            part1_sum+=int(x)

                # part 2
                if char == "*":
                    # get all adjacents
                    adj_curr_line = [line[j-1], line[j+1]]
                    adj_top_line = list(lines[i-1][j-1:j+2])
                    adj_below_line = list(lines[i+1][j-1:j+2])
                    all_adj = adj_curr_line + adj_top_line + adj_below_line  
                    
                    get_numbers = [int(x) for x in remove_adjacent_duplicates(all_adj) if x.isdigit()]
                    if len(get_numbers) == 2:
                        part2_sum += get_numbers[0] * get_numbers[1]

    print(f"Part 1: {part1_sum}")
    print(f"Part2: {part2_sum}")
                    

def calibrate(lines):
    """
    group adjacent digits to themselves, extend the list to their length.
    this way the length of the original input is not modified 
    """
    main_output = []
    for line in lines:
        line = line.strip()
        result_list = []
        matches = re.findall(r'\D|\d+', line)
        for x in matches:
            if x.isdigit():
                result_list.extend([x] * len(x))
            else:
                result_list.append(x)
        main_output.append(result_list)
    return main_output


def remove_adjacent_duplicates(lst):
    """
    removes adjacent duplicates that are adjacent symbols

    adjacent numbers are the same number due to the calibration function
    """
    result = []
    prev_element = None

    for element in lst:
        if element != prev_element:
            result.append(element)
            prev_element = element

    return result


main()
